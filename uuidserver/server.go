package uuidserver

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"uuid"
)

type UUIDHandlerFunc func() uuid.UUID

func (generator UUIDHandlerFunc) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	b, _ := generator().MarshalText()
	contentLength := len(b)
	w.Header().Set("Content-Length", strconv.Itoa(contentLength))
	w.Header().Set("Content-Type", "text/plain")
	r := bytes.NewReader(b)
	io.Copy(w, r)
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", UUIDHandlerFunc(uuid.New))
	mux.Handle("/v4", UUIDHandlerFunc(uuid.NewV4))
	mux.Handle("/v7", UUIDHandlerFunc(uuid.NewV7))
	mux.Handle("/nil", UUIDHandlerFunc(uuid.Nil))
	mux.Handle("/max", UUIDHandlerFunc(uuid.Max))
	return mux
}
