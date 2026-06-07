package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
)

type Storage interface {
	Set(key string, value string) error
	Get(key string) (string, bool)
}

type GetPayload struct {
	Key string `json:"key"`
	Val string `json:"value"`
}

type FileStore struct {
	mu       sync.RWMutex
	Pairs    map[string]string
	FileName string
}

/*
The Dependency Injection Pattern
Instead of attaching the handler to the database, we are going to create a new, separate struct just for handling the API. This new struct won't know what database it's using; it will only know that it has a Storage interface it can talk to.

Here is how you execute this.

Step 1: Create the API Controller Struct
Create a struct that holds your Storage interface. This is the core of Dependency Injection. We are "injecting" the database dependency into the API layer.
*/
// 1. This struct ONLY handles web routing. It knows nothing about files.
type APIHandler struct {
	db Storage // Notice we use the interface, not *FileStore!
}

func (f *FileStore) Set(key string, value string) error {
	fileOb, err := os.OpenFile(f.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Failed to open log file: '%s', '%w'", f.FileName, err)
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	defer fileOb.Close()
	f.Pairs[key] = value

	logLine := fmt.Sprintf(`{"key": "%s", "value": "%s"}`+"\n", key, value)

	if _, err := fileOb.Write([]byte(logLine)); err != nil {
		return fmt.Errorf("Failed to write log to file: '%s', '%w'", f.FileName, err)
	}

	return nil
}

func (f *FileStore) Get(key string) (string, bool) {
	f.mu.RLock()
	val, exist := f.Pairs[key]
	f.mu.RUnlock()
	return val, exist
}

/*
Step 2: Refactor the HTTP Handler
Now, instead of attaching your handler to FileStore, attach it to APIHandler.
*/
func (api *APIHandler) HandleSet(w http.ResponseWriter, r *http.Request) {
	var req GetPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// 3. Call the interface! The API doesn't know it's writing to a file.
	// It just trusts that `api.db` has a `Set` method.
	if err := api.db.Set(req.Key, req.Val); err != nil {
		fmt.Println("Database error: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)

}

func (api *APIHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	val, exist := api.db.Get(vars["key"])
	if !exist {
		http.Error(w, "Couldn't find value", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Key: %s has value: %s", vars["key"], val)
}

func PopulateMap(myFileStore *FileStore) error {
	file, err := os.Open(myFileStore.FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Totally fine on first run, just return
		}
		return fmt.Errorf("Failed to open file: %s, %w", myFileStore.FileName, err)
	}
	defer file.Close()

	// pairs := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		logLines := scanner.Bytes()

		var entry GetPayload

		if err := json.Unmarshal(logLines, &entry); err != nil {
			fmt.Println("Skipping unreadable line Error:", err)
			continue
		}
		myFileStore.mu.Lock()
		myFileStore.Pairs[entry.Key] = entry.Val
		myFileStore.mu.Unlock()

	}

	return nil
}

func main() {
	fmt.Println("Welcome to Persistend Key Value Store")

	// 1. Initialize your Database Layer
	myFileStore := &FileStore{
		Pairs:    make(map[string]string),
		FileName: "database.log",
	}
	// 2. Initialize your API Layer, and inject the Database!
	// Because *FileStore implements the Set and Get methods,
	// Go allows it to satisfy the Storage interface.
	apiController := &APIHandler{
		db: myFileStore,
	}

	// 3. Set up your routes using the API Controller
	r := mux.NewRouter()
	r.HandleFunc("/set", apiController.HandleSet).Methods("POST")
	r.HandleFunc("/get/{key}", apiController.HandleGet)

	PopulateMap(myFileStore)

	http.ListenAndServe(":8080", r)

}
