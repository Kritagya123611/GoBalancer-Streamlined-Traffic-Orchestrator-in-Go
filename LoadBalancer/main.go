package main

import (
    "fmt"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "time"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/cors"
)

type Backend struct {
    URL   *url.URL
    Alive bool
}

var backends []*Backend
var currentServer int

// --------- Round Robin Picker ---------
func GetNextRoute() *Backend {
    total := len(backends)
    for i := 0; i < total; i++ {
        backend := backends[currentServer]
        currentServer = (currentServer + 1) % total
        if backend.Alive {
            return backend
        }
    }
    return nil
}

// --------- Proxy Handler ---------
func HandleRoutes(w http.ResponseWriter, r *http.Request) {
    backend := GetNextRoute()
    if backend == nil {
        http.Error(w, "No healthy backends", http.StatusServiceUnavailable)
        return
    }
    proxy := httputil.NewSingleHostReverseProxy(backend.URL)
    proxy.ServeHTTP(w, r)
}

// --------- Health Checker ---------
func checkHealth(backend *Backend) {
    client := http.Client{Timeout: 2 * time.Second}
    resp, err := client.Get(backend.URL.String() + "/health")
    if err != nil || resp.StatusCode >= 500 {
        backend.Alive = false
        fmt.Println("❌ DOWN:", backend.URL)
        return
    }
    backend.Alive = true
    fmt.Println("✅ UP:", backend.URL)
}

func main() {
    rawURLs := []string{"http://localhost:8081", "http://localhost:8082"}
    for _, raw := range rawURLs {
        parsed, _ := url.Parse(raw)
        backends = append(backends, &Backend{URL: parsed, Alive: true})
    }

    // start health check loop
    go func() {
        for {
            for _, b := range backends {
                checkHealth(b)
            }
            time.Sleep(5 * time.Second)
        }
    }()

    // router
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

    fmt.Println("Load balancer running on :8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}
