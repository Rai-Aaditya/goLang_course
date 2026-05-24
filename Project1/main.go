package main

import (
	b64 "encoding/base64"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	m := make(map[string]string)
	// create a new standard library multiplexer(server)
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to URL Shortener!")
	})
	r.HandleFunc("/url/{urlVal}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		urlVal := vars["urlVal"]
		fmt.Fprintf(w, "Original url: %s\n", urlVal)
		fmt.Println("Original url:", urlVal)
	})

	r.HandleFunc("/shorten/{url}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		url := vars["url"]
		// sEnc := b64.StdEncoding.EncodeToString([]byte(url))
		sEnc := b64.URLEncoding.EncodeToString([]byte(url))
		fmt.Println("Encoded:", sEnc)
		shortenedUrl := sEnc[:6]
		m[shortenedUrl] = sEnc

		fmt.Fprintf(w, "Shortened url: %s\n", shortenedUrl)
	})

	r.HandleFunc("/redirect/{shortened}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		shortenedUrl := vars["shortened"]
		sDec, _ := b64.URLEncoding.DecodeString(m[shortenedUrl])

		http.Redirect(w, r, "http://"+string(sDec), http.StatusSeeOther)

	})

	port := "8080"
	fmt.Println("Website up and running on port:", port)
	http.ListenAndServe(":"+port, r)
}
