package main

import (
	"fmt"
	"strings"
)

type EmployeeIn interface {
	getDetails() string
	displayDetails()
}

type Employee struct {
	name  string
	age   int
	phone int
	email string
}

func (emp Employee) getDetails() {
	valid := strings.Contains(emp.email, "@")
	if !valid {
		fmt.Println("Invalid email")
	} else {
		fmt.Println("Email: ", emp.email)
	}

}

func (emp Employee) displayDetails() {
	fmt.Println("Name: ", emp.name)
	fmt.Println("Age: ", emp.age)
	fmt.Println("Phone: ", emp.phone)

}

func main() {
	em := Employee{"John", 30, 1234567890, "asdf@comp.com"}
	em.getDetails()
	em.displayDetails()
}
