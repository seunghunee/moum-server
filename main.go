package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc(pathPrefix, apiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

const pathPrefix = "/api/v1/article/"
