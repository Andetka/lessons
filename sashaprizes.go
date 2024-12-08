package main

import (
	"fmt"
)

// у нас есть мапа с курсами валют
// ЕВРО - 120 рублей
// ДОЛЛАР - 100 рублей
//
// и есть список призов (250к рублей, 1кк рублей, 30к рублей).
//
// Нужно вывести на экран сколько в евро каждый приз
//
// и после отдельно вывести сколько в долларах каждый приз
//
// т.е я хочу ввидеть типа
//
// приз 1 = 2500 евро
// приз 2 10000 евро
// приз 30 300 евро
//
// приз 1 = 2500 долларов
// приз 2 10000 долларов
// приз 30 300 долларов

func main() {
	moneys := map[string]int{"euro": 120, "dollar": 100}
	prizes := map[string]int{"Приз 1 -": 250000, "Приз 2 -": 1000000, "Приз 3 -": 30000}
	for prizeNumb, prize := range prizes {
		euro := prize / moneys["euro"]
		fmt.Println(prizeNumb, euro, "евро")
	}

	for prizeNumb, prize := range prizes {
		dollar := prize / moneys["dollar"]
		fmt.Println(prizeNumb, dollar, "долларов")
	}
}
