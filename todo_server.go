package main

import (
	"fmt"
	"log"
	"net/http"
)

func todoHandler(w http.ResponseWriter, req *http.Request) {
	method := "GET"
	fmt.Fprintf(w, "hello %s todos", method)
}

func main() {
	http.HandleFunc("/todos", todoHandler)

	log.Fatal(http.ListenAndServe(":1234", nil))
}
