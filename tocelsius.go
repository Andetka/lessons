package main

import (
	"fmt"
	"log"

	"awesomeProject/keyboard"
)

func main() {
	fmt.Print("Введи температуру в Фарингейтах:")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f градусов по Цельсия\n", celsius)
}
