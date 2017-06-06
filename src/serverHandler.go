package main

import (
	"net/http"
	"handler"
)

func main() {
	base := new(handler.Base)
	http.Handle("/", base)
	http.ListenAndServe(":8083", nil)
}
