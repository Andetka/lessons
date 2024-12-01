package main

import (
	"fmt"
)

func main() {
	var notes [7]string
	notes[0] = "do"
	fmt.Println(notes[3]) // здесь будет нулевое значение
	fmt.Println(notes[6]) // здесь будет нулевое значение
	fmt.Println(notes[0])

	var primes [5]int
	primes[0] = 2
	fmt.Println(primes[0])
	fmt.Println(primes[2]) // здесь будет нулевое значение
	fmt.Println(primes[4]) // здесь будет нулевое значение

	var notes2 = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"} // это то же самое, но записано через литерал массивов
	fmt.Println(notes2[3], notes2[6], notes2[0])

	var primes2 = [5]int{2, 3, 5, 7, 11} // это то же самое, но записано через литерал массивов
	fmt.Println(primes2[0], primes2[2], primes2[4])

	notes3 := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"} // это то же самое, но записано через короткое объявление
	primes3 := [5]int{2, 3, 5, 7, 11}                             // это то же самое, но записано через короткое объявление
	fmt.Println(notes3[5], primes3[3])
}
