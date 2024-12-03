package main

import (
	"fmt"
)

func mix(flag bool, num int, strings ...string) {
	fmt.Println(num, flag, strings)
}
func main() {
	mix(true, 1, "a", "b")
	mix(false, 2, "a", "b", "c", "d")
}
