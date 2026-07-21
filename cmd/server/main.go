package main

import (
	"log"
	"net/http"

	delivery "github.com/Devesanoff/symbolix.uz/internal/delivery/http"
)

func main() {
	mux := http.NewServeMux()

	// API routes
	mux.Handle("/api/v1/ping", delivery.CORS(http.HandlerFunc(delivery.PingHandler)))
	mux.Handle("/api/v1/convert/krill-lotin", delivery.CORS(http.HandlerFunc(delivery.CyrillicToLatinHandler)))
	mux.Handle("/api/v1/convert/lotin-krill", delivery.CORS(http.HandlerFunc(delivery.LatinToCyrillicHandler)))
	mux.Handle("/api/v1/convert/pdf-docx", delivery.CORS(http.HandlerFunc(delivery.PDFToDocxHandler)))

	// Static frontend — web/ papkasidan xizmat ko'rsatish
	fs := http.FileServer(http.Dir("./web"))
	mux.Handle("/", fs)

	addr := ":8080"
	log.Printf("Server starting on %s", addr)
	log.Printf("Frontend available at http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}