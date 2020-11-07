package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	resp := []byte(`{"name": "anuchit"}`)
	w.Write(resp)
}

func main() {
	fmt.Println("start...")
	http.HandleFunc("/", helloHandler)

	fmt.Println("end....")
}
