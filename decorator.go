package main

import (
	"log"
	"net/http"
	"time"
)

// Logger = "first"-"outer"-http-handler. Logger = decorator of "nextHandler"-a.
// input arg "nextHandler" = "second"-"inner"-http-handler.
func Logger(nextHandler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // start wrap(decorate) by "outer"-logger-handler:
        start := time.Now()

        // call "inner"-http-handler.
        nextHandler.ServeHTTP(w, r)

        // complete wrap(decorate) by "outer"-logger-handler:
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}

func Main_decorator() {
    mux := http.NewServeMux()
    
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })

    // Decorate the entire mux with the Logger.
    http.ListenAndServe(":8080", Logger(mux))
}







