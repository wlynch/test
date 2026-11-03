package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler responds to every request with a greeting.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	addr := ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
