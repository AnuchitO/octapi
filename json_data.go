package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Birthday struct {
	time.Time
}

func (bb *Birthday) UnmarshalJSON(b []byte) error {
	fmt.Println(string(b))
	s := strings.Trim(string(b), "\"")
	fmt.Println(s)

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	bb.Time = t
	return nil
}

type Person struct {
	Birthday Birthday `json:"birthday"`
}

func main() {
	payload := bytes.NewBufferString(`{ "birthday":"1989-02-14" }`)

	p := Person{}

	err := json.Unmarshal(payload.Bytes(), &p)
	fmt.Println(err)
	fmt.Println(p)
	fmt.Println(p.Birthday)
}
