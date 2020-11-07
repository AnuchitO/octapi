package main

import (
	"fmt"
	"log"
	"net/http"
)

func todoHandler(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	if method == "POST" {
		fmt.Fprintf(w, "hello %s created todos", method)
		return
	}

	fmt.Fprintf(w, "hello %s todos", method)
}

func main() {
	http.HandleFunc("/todos", todoHandler)

	log.Fatal(http.ListenAndServe(":1234", nil))
}
