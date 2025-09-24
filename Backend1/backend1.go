package main

import (
	"fmt"
	"net/http"
)

var count int

func metricsBackend1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"requests": %d}`, count)
}

func backend1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("backend1 got a request")
	count++
	fmt.Println("Hello from Backend 1 — hit %d times\n",count)
}

func main() {
	http.HandleFunc("/", backend1)
	http.HandleFunc("/metric1",metricsBackend1)
	fmt.Println("Backend1 running on :8081")
	http.ListenAndServe(":8081", nil)
}
