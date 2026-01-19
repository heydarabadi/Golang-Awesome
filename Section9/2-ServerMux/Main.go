package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type TestHandler struct {
}

func (h TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)
	fmt.Fprintf(w, "hello world")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Pouya,Ali,Rezvan,Hossein,Mohammad")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.RedirectHandler("https://google.com", 307))
	mux.Handle("/yahoo", http.RedirectHandler("https://yahoo.com", 307))
	mux.HandleFunc("/Get-Users", GetUsers)

	server1 := &http.Server{Addr: ":8091", Handler: mux /*TestHandler{} */, WriteTimeout: 5 * time.Second, ReadTimeout: 5 * time.Second}

	err := server1.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
