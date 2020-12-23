package main

import (
	"fmt"
	"github.com/explodes/explodio/stand"
	"net/http"
)

func errorMiddleware(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			fmt.Printf("ERROR %s: %s", r.URL, err)
		}
	}
}

func health(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(200)
	if _, err := w.Write([]byte("ok")); err != nil {
		return err
	}
	return nil
}

func hello(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(200)
	if _, err := w.Write([]byte("hello")); err != nil {
		return err
	}
	return nil
}

func main() {
	http.HandleFunc("/health", errorMiddleware(health))
	http.HandleFunc("/", errorMiddleware(hello))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", stand.RequireEnv("HTTP_PORT")), nil); err != nil {
		panic(err)
	}
}
