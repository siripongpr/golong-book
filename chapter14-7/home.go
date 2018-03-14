package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomePageHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if vars["name"] == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func UsersHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Users Page")
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/{name}", HomePageHandle).Methods("GET")
	r.HandleFunc("/users", UsersHandle).Methods("GET")
	return r
}

func main() {
	var port string

	flag.StringVar(&port, "port", ":3000", "default port: 3000")
	flag.Parse()

	http.ListenAndServe(port, NewRouter())
}

/*
go build -o main.exe // for windows
go build -o main // for linux, mac
>./main.exe -port=:8080 // for windows
>./main -port=:8080
*/
