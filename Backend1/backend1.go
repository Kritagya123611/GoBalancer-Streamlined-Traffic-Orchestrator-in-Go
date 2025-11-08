package main

import (
	"fmt"
	"net/http"
	//"github.com/go-chi/chi/v5"
)

var count int

func metricsBackend1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"requests": %d}`, count)
}
//testing comment
//testing comment
//another testing comment
//yet another testing comment
//one more testing comment
//final testing comment
func backend1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("backend1 got a request")
	count++
	fmt.Printf("Hello from Backend 1 — hit %d times\n",count)
}

func healthCheck(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)                     
	fmt.Fprintln(w, `{"status": "healthy"}`)	
}

func main() {
	http.HandleFunc("/", backend1)
	http.HandleFunc("/health",healthCheck)
	http.HandleFunc("/metric1",metricsBackend1)
	fmt.Println("Backend1 running on :8081")
	http.ListenAndServe(":8081", nil)
}
