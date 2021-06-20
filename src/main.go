package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

type MyHandler struct{}

func (h *MyHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server{
		Addr:		"127.0.0.1:8000",
	}

	http.Handle("/", &handler)
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)

	server.ListenAndServe()

	// https用 設定
	// cert.pem :SSL証明書
	// key.pem	:サーバ証明書
	// server.ListenAndServeTLS("cert.pem", "key.pem")
}