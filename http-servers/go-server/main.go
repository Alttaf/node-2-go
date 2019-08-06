package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func handler(w http.ResponseWriter, r *http.Request) {

	for i := 0; i < 2000000; i++ {
		uuid.New()
		fmt.Fprintf(w, "%d", i)
	}

}

func main() {
	port := ":8080"

	http.HandleFunc("/", handler)
	fmt.Print("starting http server listening on port" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
