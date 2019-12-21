package main

import (
	"fmt"
)

var (
	example = []string{
		"golang",
		"hands-on",
		"in",
		"umeda",
	}
)

func main() {
	var data []string
	data = example
	for _, v := range data {
		fmt.Print(v)

		if v == "in" {
			fmt.Println(" ○")
		} else {
			fmt.Println(" ✗")
		}
	}
}
