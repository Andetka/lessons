package main

import (
	"fmt"
)

func main() {
	names := []string{"sasha", "nastya", "karina", "sasha", "nastya"}
	uniqueNamesCount := make(map[string]int)
	for _, name := range names {
	
		uniqueNamesCount[name]++
	}
	for name, count := range uniqueNamesCount {
		fmt.Printf("%s:%d\n", name, count)
	}

}
