package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	resp := []byte(`{"name": "anuchit"}`)
	w.Write(resp)
}
// http://localhost:1234/
func main() {
	fmt.Println("start...")
	http.HandleFunc("/", helloHandler)

	http.ListenAndServe(":1234567890", nil) // ":1234567890"
	fmt.Println("end....")
}
