package main

import (
	"log"
	"net/http"

	"github.com/faruqii/Go-GraphQL/handler"
)

func main() {
	http.HandleFunc("/graphql", handler.GraphqlHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}