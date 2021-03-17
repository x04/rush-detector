package main

import (
	"log"
	"net/http"
	"rushDetector/http2"
)

func main() {
	var srv http.Server
	srv.Addr = "0.0.0.0:443"
	http2.ConfigureServer(&srv, &http2.Server{})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello not rush client!"))
	})

	go func() {
		log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
	}()
	select {}
}
