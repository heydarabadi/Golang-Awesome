package main

import (
	"fmt"
	"log"
	"net/http"
)

type UserHandler struct {
}

type user struct {
	Id   int
	Name string
}

func (u UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && len(r.URL.Query().Get("id")) == 1:
		GetUser(w, r)
		return

	case r.Method == http.MethodGet && len(r.URL.Query().Get("id")) == 0:
		GetUsers(w, r)
		return
	}

}

func CreateUser(u user, w http.ResponseWriter, r *http.Request) {
	getData := r.Header.Get("User-Name")
	fmt.Fprintf(w, getData)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get User By Id")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get All Users")

}

func main() {

	errForCreateServer := http.ListenAndServe(":8090", UserHandler{})

	if errForCreateServer != nil {
		log.Fatal(errForCreateServer)
	}

}
