package main

import (
	"fmt"

	"awesomeProject/magazine"
)

func printInfoSub(s *magazine.Subscriber) {
	fmt.Println("Имя:", s.Name, "\nРейтинг:", s.Rate, "\nАктивность:", s.Active, "\nАдрес:", s.Address)
}
func printInfoEmp(s *magazine.Employee) {
	fmt.Println("Имя:", s.Name, "\nЗарплата:", s.Salary, "$", s.Address)
}
func defaultSubscriber(name string) *magazine.Subscriber {
	var s magazine.Subscriber
	s.Name = name
	s.Rate = 5.0
	s.Active = true
	return &s
}

func applyDiscount(s *magazine.Subscriber) {
	s.Rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("Карина")
	applyDiscount(subscriber1)
	subscriber1.City = "Москва"              // subscriber1.Address.City = "Москва"
	subscriber1.Country = "Россия"           // subscriber1.Address.Country
	subscriber1.PostalCode = "127411"        // subscriber1.Address.PostalCode
	subscriber1.Street = "Дмитровское шоссе" // subscriber1.Address.Street
	printInfoSub(subscriber1)

	subscriber2 := defaultSubscriber("Александр")
	subscriber2.Rate = 4.7
	printInfoSub(subscriber2)

	subscriber4 := magazine.Subscriber{Name: "Артём", Rate: 5.0, Active: true}
	printInfoSub(&subscriber4)

	subscriber5 := magazine.Subscriber{Name: "Настя"}
	printInfoSub(&subscriber5)

	employee1 := magazine.Employee{Name: "Аня", Salary: 1000}
	employee1.City = "Москва"
	employee1.Country = "Россия"
	employee1.PostalCode = "123456"
	employee1.Street = "Долгопрудный"
	printInfoEmp(&employee1)
}
