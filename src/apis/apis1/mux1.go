package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().PathPrefix("/test").Methods("GET").Subrouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products/{productId}", ProductsHandler)
	r = r.NewRoute().Methods("POST").PathPrefix("/test1").Headers("Content-Type", "application/(text|json)").Subrouter()
	r.HandleFunc("/candidate", InsertCandidate)
	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

//HomeHandler home handler
func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("HomeHandler")
	//return "HomeHandler", nil
}

//ProductsHandler home handler
func ProductsHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("ProductsHandler")
	vars := mux.Vars(request)
	productId := vars["productId"]
	fmt.Println(productId)
	writer.Write([]byte(productId))
	//return "HomeHandler", nil
}

//ArticlesHandler ArticlesHandler
func InsertCandidate(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var c Candidate
	decoder.Decode(&c)
	fmt.Println(c)
	encoder := json.NewEncoder(writer)
	encoder.Encode(&c)
	//return "HomeHandler", nil
}

//Candidate
type Candidate struct {
	Name  string
	Email string
	age   int
}
