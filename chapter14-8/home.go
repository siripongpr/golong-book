package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func HomePageHandle(w http.ResponseWriter, r *http.Request) {
	//start := time.Now()
	//fmt.Printf("Start at %v\n", start)
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
	//fmt.Printf("Completed in %v\n", time.Since(start))
}

func UsersHandle(w http.ResponseWriter, r *http.Request) {
	//start := time.Now()
	//fmt.Printf("Start at %v\n", start)
	fmt.Fprintf(w, "Users Page")
	//fmt.Printf("Completed in %v\n", time.Since(start))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Start at %v\n", start)

		next.ServeHTTP(w, r)

		fmt.Printf("Completed in %v\n", time.Since(start))
	})
}
func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", HomePageHandle).Methods("GET")
	r.HandleFunc("/users", UsersHandle).Methods("GET")
	r.Use(loggingMiddleware)
	return r
}

func main() {
	http.ListenAndServe(":3000", NewRouter())
}

/* If do not use gorilla/mux
func loggingMiddleware(next func(http.ResponseWriter, *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Start")

		next(w, r)

		fmt.Printf("Completed in %v", time.Since(start))
	}
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", loggingMiddleware(HomePageHandle)).Methods("GET")
	r.HandleFunc("/users", loggingMiddleware(UsersHandle)).Methods("GET")
	//r.Use(loggingMiddleware)
	return r
}
*/
