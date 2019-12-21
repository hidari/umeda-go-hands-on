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
		fmt.Println(v)

		fmt.Print("input: ")
		var ans string
		fmt.Scanln(&ans)

		if v == ans {
			fmt.Println("o")
		} else {
			fmt.Println("x")
		}
	}
}
