package main

import (
	"handler"
	"net"
	"net/http"
)

func main() {
	base := new(handler.Base)
	http.Handle("/", base)
	http.ListenAndServe(":8083", nil)
	net.Dial()
	net.Listen("http", "localhost:8083")
}
