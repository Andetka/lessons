package main

import (
	"fmt"
)

func main() {

	ranks := make(map[string]int)
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["gold"])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "Lithium"
	fmt.Println(elements["H"])
	fmt.Println(elements["Li"])

	isPrime := make(map[int]bool)
	isPrime[4] = false
	isPrime[7] = true
	fmt.Println(isPrime[4])
	fmt.Println(isPrime[7])

	jewelry := make(map[string]float64)
	jewelry["necklace"] = 89.99
	jewelry["earrings"] = 89.99
	clothing := map[string]float64{"pants": 59.99, "shirt": 39.99}
	fmt.Println("Earrings:", jewelry["earrings"])
	fmt.Println("Necklace", jewelry["necklace"])
	fmt.Println("Shirt", clothing["shirt"])
	fmt.Println("Pants:", clothing["pants"])
}
