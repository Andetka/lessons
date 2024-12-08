package main

import (
	"fmt"
	"log"
	"strconv"
)

// есть массив в котором хранятся пары имя и возраст [sasha 28 karina 29 nasta 27 sasha 28]
// Нужно избавиться от дубликатов и выести итоговый список име
// sasha 28
// nasta 27
// karina 29

func main() {
	slice := []string{"sasha", "28", "karina", "29", "nastya", "27", "sasha", "28"}
	list := make(map[string]int)

	for i := 0; i < len(slice); i = i + 2 {
		key := slice[i]
		value, err := strconv.Atoi(slice[i+1])
		if err != nil {
			log.Fatal(err)
		}
		list[key] = value

	}
	for key, value := range list {
		fmt.Println(key, value)
	}

}
