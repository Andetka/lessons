// программа pass_fail сообщает, сдал ли пользователь экзамен

package main

import (
	"fmt"
	"log"

	"awesomeProject/keyboard"
)

func getStatus(grade float64) string {
	if grade >= 60 {
		return "passing"
	}

	return "failing"
}

func main() {
	fmt.Print("Enter a grade:")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	status := getStatus(grade)
	fmt.Println("A grade of", grade, "is", status)

}
