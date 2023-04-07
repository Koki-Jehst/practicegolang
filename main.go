package main

import (
	"fmt"
	"net/http"
)

func h(web http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(web, "welcome to my test server!")
	fmt.Println(request)
}

func main() {
	http.HandleFunc("/", h)
	http.ListenAndServe(":8080", nil)
}