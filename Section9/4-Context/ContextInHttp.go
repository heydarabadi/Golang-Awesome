package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type RequestHandler struct {
}

func (h RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-ctx.Done():
		println("Request Cancelled")
		return

	case <-time.After(time.Second * 5):
		_, err := fmt.Fprintf(w, "Processing Request")
		if err != nil {
			return
		}
		println("Processing Start")
		return
	}
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", RequestHandler{})

	errForCreateServer := http.ListenAndServe(":8090", RequestHandler{})

	if errForCreateServer == nil {
		log.Fatal("Error creating server ", mux)
	}

}
