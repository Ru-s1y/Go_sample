package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr:		"127.0.0.1:8000",
		Handler:	nil,
	}
	server.ListenAndServe()

	// https用 設定
	// cert.pem :SSL証明書
	// key.pem	:サーバ証明書
	// server.ListenAndServeTLS("cert.pem", "key.pem")
}