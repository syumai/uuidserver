package uuidserver

import (
	"bytes"
	"io"
	"net/http"
	"uuid"
)

func newUUIDReader(generator func() uuid.UUID) io.Reader {
	b, _ := generator().MarshalText()
	return bytes.NewReader(b)
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, newUUIDReader(uuid.New))
	})
	mux.HandleFunc("/v4", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, newUUIDReader(uuid.NewV4))
	})
	mux.HandleFunc("/v7", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, newUUIDReader(uuid.NewV7))
	})
	mux.HandleFunc("/nil", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, newUUIDReader(uuid.Nil))
	})
	mux.HandleFunc("/max", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, newUUIDReader(uuid.Max))
	})
	return mux
}
