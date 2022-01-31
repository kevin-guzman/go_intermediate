package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Printer interface {
	getMessage() string
}

type Employee struct {
	id int
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "FT message"
}

func (ttEmployee TemporaryEmployee) getMessage() string {
	return "TT message"
}

func getMessage(p Printer) {
	fmt.Println("Message->", p.getMessage())
}

func main() {
	fe := FullTimeEmployee{}
	fe.age = 12
	fe.id = 1
	fe.name = "Kevin"
	getMessage(fe)
	te := TemporaryEmployee{}
	getMessage(te)
}
