// Writing a code to shorten url: using map to store the pair of shortened url and original url, previous code was working, but to make it ready for production and more microservice friendly, writing this code understanding instructions from gemini:
// Chat url: https://gemini.google.com/share/f2a145d164a1

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func getShortenCode(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	shortCode := make([]byte, length)

	for i := range shortCode {
		shortCode[i] = charset[r.Intn(len(charset))]
	}

	return string(shortCode)

}

// 1. define your payload structure using struct tags
type ShortenRequest struct {
	URL string `json:"url"`
}

// 2. Encapsulate your data store (This is where mutex will go)
type URLStore struct {
	mu   sync.RWMutex
	urls map[string]string
}

// 3. Attach your handler as a method to the store
func (s *URLStore) HandleShorten(w http.ResponseWriter, r *http.Request) {
	// we expect a POST request with a JSON body: {"url": "https://example.com"}
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		req.URL = "https://" + req.URL
	}
	// TODO: Generate a random 6-character string
	shortcode := getShortenCode(6)
	// Save to map
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urls[shortcode] = req.URL

	// Send the JSON Response back
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"short_url": "%s"}`, shortcode)

}

func (s *URLStore) RedirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	shortCode := vars["shortcode"]

	s.mu.RLock()

	originalURL, exists := s.urls[shortCode]

	s.mu.RUnlock()

	if !exists {
		http.Error(w, "URL not found!", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusSeeOther)

}

func main() {
	// Initiate your struct
	store := &URLStore{
		urls: make(map[string]string),
	}

	r := mux.NewRouter()
	// bg := &sync.WaitGroup{}	// another way to declare struct
	r.HandleFunc("/shorten", store.HandleShorten).Methods("POST")
	r.HandleFunc("/{shortcode}", store.RedirectURL)
	// Pass the method reference to the router

	port := "8080"

	fmt.Printf("Server running on: %s\n", port)
	http.ListenAndServe(":"+port, r)
}
