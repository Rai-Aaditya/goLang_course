// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// 	"time"

// 	"github.com/gorilla/mux"
// )

// type Client struct {
// 	count    int
// 	lastSeen time.Time
// }

// type MemoryMap struct {
// 	mu sync.Mutex
// 	mp map[string]*Client
// }

// func getDataHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, `{"message": "Success! data fetched"}`)
// }

// func main() {

// 	initMap := &MemoryMap{
// 		mp: make(map[string]*Client),
// 	}

// 	fmt.Println("API Rate Limiter")

// 	r := mux.NewRouter()
// 	r.HandleFunc("/api/data", getDataHandler).Methods("GET")

// 	http.ListenAndServe(":8080", r)
// }
