package main

import (
	"fmt"
	"log"
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

const (
	ip   = "127.0.0.1" // should be empty for prod deploys
	port = "8080"
)

func main() {
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/invert", invertHandler)
	http.HandleFunc("/flatten", flattenHandler)
	http.HandleFunc("/sum", sumHandler)
	http.HandleFunc("/multiply", multiplyHandler)

	addr := fmt.Sprintf("%s:%s", ip, port)

	log.Printf("Starting server at %s...\n", addr)

	log.Fatal("Starting server", http.ListenAndServe(addr, nil))
}
