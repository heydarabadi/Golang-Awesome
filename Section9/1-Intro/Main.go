package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	// Method 1
	var server1 = &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	err := server1.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	// Method 2
	go func() {
		http.ListenAndServe(":8090", nil)
	}()

}
