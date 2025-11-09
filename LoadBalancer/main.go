package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Backend struct {
	URL    *url.URL
	Alive  atomic.Bool
	Weight int
}

var backends []*Backend
var current int32

func buildWeightedPool() []int {
	var pool []int
	for i, b := range backends {
		if b.Alive.Load() {
			for j := 0; j < b.Weight; j++ {
				pool = append(pool, i)
			}
		}
	}
	return pool
}

func GetNextRoute() *Backend {
	weightedPool := buildWeightedPool()
	if len(weightedPool) == 0 {
		return nil
	}
	idx := atomic.AddInt32(&current, 1) % int32(len(weightedPool))
	return backends[weightedPool[idx]]
}

func HandleRoutes(w http.ResponseWriter, r *http.Request) {
	backend := GetNextRoute()
	if backend == nil {
		http.Error(w, "No healthy backends", http.StatusServiceUnavailable)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(backend.URL)
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		http.Error(w, "Backend error", http.StatusBadGateway)
	}
	proxy.ServeHTTP(w, r)
}

func checkHealth(backend *Backend) {
	client := http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get(backend.URL.String() + "/health")
	if err != nil || resp.StatusCode >= 500 {
		if backend.Alive.Swap(false) {
			fmt.Println("DOWN:", backend.URL)
		}
		return
	}
	if !backend.Alive.Swap(true) {
		fmt.Println("UP:", backend.URL)
	}
}

func main() {
	rawURLs := []struct {
		URL    string
		Weight int
	}{
		{"http://localhost:8081", 3}, 
		{"http://localhost:8082", 1}, 
	}

	for _, raw := range rawURLs {
		parsed, _ := url.Parse(raw.URL)
		b := &Backend{URL: parsed, Weight: raw.Weight}
		b.Alive.Store(true) 
		backends = append(backends, b)
	}

	go func() {
		for {
			for _, b := range backends {
				checkHealth(b)
			}
			time.Sleep(5 * time.Second)
		}
	}()
// Setting up router with CORS
//testing logs failure
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	router.HandleFunc("/*", HandleRoutes)

	srv := &http.Server{Addr: ":8000", Handler: router}

	go func() {
		fmt.Println("Load balancer running on :8000")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
