package main

import (
	"fmt"
	"log"
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

	log.Fatal(http.ListenAndServe(":1234", nil))
	fmt.Println("end....")
}
