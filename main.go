package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"
)

//go:embed static/index.html
var content embed.FS

func main() {
	port := flag.Int("port", 9090, "listen port")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := content.ReadFile("static/index.html")
		if err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(data)
	})

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Listening on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}