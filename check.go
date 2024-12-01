package main

import (
	"fmt"
)

func main() {
	var slice []string
	if len(slice) == 0 {
		slice = append(slice, "first item")
	}
	fmt.Println(slice)
}
