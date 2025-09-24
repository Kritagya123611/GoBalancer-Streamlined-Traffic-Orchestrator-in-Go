package main

import (
	"fmt"
	"net/http"
)

func backend2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("backend2 got a request")
	fmt.Fprintln(w, "Hello from Backend 2")
}

func main() {
	http.HandleFunc("/", backend2)
	fmt.Println("Backend2 running on :8082")
	http.ListenAndServe(":8082", nil)
}
