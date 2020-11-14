package main

import (
	"fmt"
	"github.com/anuchito/octapi/a"
	"github.com/anuchito/octapi/b"
)

func init() {
	fmt.Println("package : main init")
}

func main() {
	b.Show()
	a.Show()
}
