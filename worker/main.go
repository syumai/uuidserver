package main

import (
	"github.com/syumai/uuidserver/uuidserver"
	"github.com/syumai/workers"
)

func main() {
	h := uuidserver.NewHandler()
	workers.Serve(h) // use http.DefaultServeMux
}
