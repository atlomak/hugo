package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ZeroReader struct{}

func (z ZeroReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 0
	}
	return len(p), nil
}

func parseSize(s string) int64 {
	s = strings.ToUpper(s)
	unit := int64(1)
	if strings.HasSuffix(s, "G") {
		unit = 1024 * 1024 * 1024
		s = s[:len(s)-1]
	}
	if strings.HasSuffix(s, "M") {
		unit = 1024 * 1024
		s = s[:len(s)-1]
	}
	if strings.HasSuffix(s, "K") {
		unit = 1024
		s = s[:len(s)-1]
	}

	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return -1
	}
	return val * unit
}

func handler(w http.ResponseWriter, r *http.Request) {
	sizeStr := strings.TrimPrefix(r.URL.Path, "/")
	size := parseSize(sizeStr)

	if size < 0 {
		http.Error(w, "Invalid size. Use /100, /10K, /50M, /1G", http.StatusBadRequest)
		return
	}
	log.Printf("Client IP: %s | Serving %s bytes", r.RemoteAddr, sizeStr)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.FormatInt(size, 10))
	w.WriteHeader(http.StatusOK)

	io.CopyN(w, ZeroReader{}, size)
}

func main() {
	port := flag.String("p", "8000", "port to listen on")
	host := flag.String("i", "", "interface to listen on")
	flag.Parse()

	addr := *host + ":" + *port
	http.HandleFunc("/", handler)
	log.Printf("Server running on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
