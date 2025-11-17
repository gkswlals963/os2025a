package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func main() {
	var s1 magazine.Subscriber
	var e1 magazine.Employee
	var e2 magazine.Employee
	var s2 magazine.Subscriber
	s1.Name = "Choi Inha"
	s2.Name = "Lee Inha"
	e1.Salary = 5000000
	e2.Salary = 5000000
	s1.Address.City = "Incheon"
	s2.Address.City = "Seoul"
	fmt.Println(s1.Name, e1.Salary, s1.Address.City)
	fmt.Println(s2.Name, e2.Salary, s2.Address.City)
}
