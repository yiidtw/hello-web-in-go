package main

import (
	"fmt"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Web")
}

func main() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
