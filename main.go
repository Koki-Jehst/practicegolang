package main

import (
	"fmt"
	"net/http"
)

func h(web http.ResponseWriter, Request *http.Request) {
	fmt.Fprintf(web, "welcome to my test server!")
}

func main() {
	http.HandleFunc("/", h)
	http.ListenAndServe(":8080", nil)
}