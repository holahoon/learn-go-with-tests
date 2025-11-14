package main

import (
	"log"
	"net/http"

	dependencyinjection "github.com/holahoon/learn-go-with-tests/8_dependency_injection"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependencyinjection.MyGreeterHandler)))
}
