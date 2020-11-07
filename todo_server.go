package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func todoHandler(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	if method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
			return
		}

		t := Todo{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
			return
		}

		fmt.Printf("body : % #v\n", t)

		fmt.Fprintf(w, "hello %s created todos", method)
		return
	}

	fmt.Fprintf(w, "hello %s todos", method)
}

func main() {
	http.HandleFunc("/todos", todoHandler)

	log.Fatal(http.ListenAndServe(":1234", nil))
}
