package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go World")
}

func main() {
	port := ":8081"
	http.HandleFunc("/", handler)
	fmt.Printf("Starting a Go http server listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
