package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world\n")

}

func main() {
	port := ":8080"

	http.HandleFunc("/", handler)
	fmt.Print("starting http server listening on port" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
