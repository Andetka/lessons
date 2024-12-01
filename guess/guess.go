// guess - игра, в которой игрок должен угадать случайное число.

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		target := rand.Intn(100) + 1
		fmt.Println("Я выбрал случайное число от 1 до 100")
		fmt.Println("Сможешь отгадать его?")

		reader := bufio.NewReader(os.Stdin)

		for guesses := 10; guesses > 0; guesses-- {
			fmt.Println("Попыток осталось:", guesses)

			fmt.Print("Введи число:")
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			guess, err := strconv.Atoi(strings.TrimSpace(input))
			if err != nil {
				log.Fatal(err)
			}

			if guess < target {
				fmt.Println("Бери выше")
			} else if guess > target {
				fmt.Println("Моё число меньше")
			} else {
				fmt.Println("МОЛОДЕЦ! ТЫ УГАДАЛ!")
				fmt.Println()
				break
			}

			if guesses == 1 {
				fmt.Println("Ты так и не смог угадать моё число. Я загадал число", target)
				fmt.Println()
			}
		}

	}
}
