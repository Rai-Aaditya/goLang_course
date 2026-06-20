package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Client struct {
	count    int
	lastSeen time.Time
}

func (app *Application) getDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Success! data fetched"}`)
}

type Application struct {
	mu sync.Mutex
	mp map[string]*Client
}

func (app *Application) routes() http.Handler {
	goMux := mux.NewRouter()
	goMux.HandleFunc("/api/data", app.getDataHandler).Methods("GET")
	return app.rateLimit(commonHeaders(goMux))
}

func commonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "Go")
		next.ServeHTTP(w, r)
	})
}

func (app *Application) rateLimit(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ipAddr, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ipAddr = r.RemoteAddr
		}

		app.mu.Lock()
		cl, exist := app.mp[ipAddr]
		if !exist {
			cl = &Client{
				count:    1,
				lastSeen: time.Now(),
			}
			app.mp[ipAddr] = cl
		} else {
			if time.Since(cl.lastSeen) > 60*time.Second {
				cl.count = 1
				cl.lastSeen = time.Now()
			} else if cl.count >= 5 {
				app.mu.Unlock()
				http.Error(w, "Rate Limit Exceeded", http.StatusTooManyRequests)
				return
			} else {
				cl.count++
			}
		}

		app.mu.Unlock()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (app *Application) Sweeper() {
	for {
		app.mu.Lock()
		for ip, cl := range app.mp {
			if time.Since(cl.lastSeen) > 3*time.Minute {
				delete(app.mp, ip)
			}
		}
		app.mu.Unlock()
		time.Sleep(3 * time.Minute)
	}
}

func main() {

	app := &Application{
		mp: make(map[string]*Client),
	}

	go app.Sweeper()
	http.ListenAndServe(":8080", app.routes())
}
