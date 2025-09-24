// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil" // ✅ for Reverse Proxy
	"net/url"           // ✅ for parsing backend URL
	"github.com/go-chi/chi/v5"   // ✅ chi router (v5 is the latest)
	"github.com/go-chi/cors"     // ✅ CORS middleware
)

var backends=[]string{
	"http://localhost:8081",
    "http://localhost:8082",
}

var CurrentServer int 

func GetNextRoute() string {
    target := backends[CurrentServer]
    CurrentServer = (CurrentServer + 1) % len(backends)
    return target
}
//backend1==server1
//backend2==server2
//ispe jo request aaye wo aaye 8000 se aur uske baad ye usko dusro mein distribute karde 
//chahe wo backend 1 ho ya phir backend 2

func HandleRoutes(w http.ResponseWriter, r *http.Request){
	target:=GetNextRoute()
	targetURL, err := url.Parse(target)
	if err != nil {
    http.Error(w, "Invalid backend", http.StatusInternalServerError)
    return
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.ServeHTTP(w, r)

}

func main(){
	router:=chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	router.HandleFunc("/*", HandleRoutes)

    fmt.Println("Load balancer running on :8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}