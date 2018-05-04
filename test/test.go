package main

import (
	"fmt"
)

type templateParams struct {
        Notice string
        Name   string
}

func main() {
	t := templateParams{}
	fmt.Println(t.Notice)
}
