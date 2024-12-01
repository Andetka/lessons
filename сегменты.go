package main

import (
	"fmt"
)

func main() {
	notes := make([]string, 7) // можно создать так
	primes := make([]int, 5)   // можно создать так
	fmt.Println(len(notes))
	fmt.Println(len(primes))

	letters := []string{"a", "b", "c"} // а можно создать так
	// а можно добавить еще элементов, которых не было при создании сегмента
	letters = append(letters, "d")
	letters = append(letters, "e")
	// и можно добавить элемент в существующий слайс, но записать новый слайс в новую переменную
	newletters := append(letters, "f")

	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}
	for _, letter := range letters {
		fmt.Println(letter)
	}

	fmt.Println(newletters)

}
