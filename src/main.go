package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function colled - " + name)
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr:		"127.0.0.1:8000",
	}

	http.HandleFunc("/", log(home))
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	// http用
	server.ListenAndServe()

	// https用 設定
	// cert.pem :SSL証明書
	// key.pem	:サーバ証明書
	// server.ListenAndServeTLS("cert.pem", "key.pem")
}