package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e Employee) GetId() int {
	return e.id
}

func (e Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}
	fmt.Println("Employee without id, name ", e)
	e.SetId(20)
	e.SetName("Wey")
	fmt.Println("Employee with setted values ", e)
	fmt.Println("Getted name by method ", e.GetName())
	fmt.Println("Getted id by method ", e.GetId())
	// new - returns pointer
	e2 := new(Employee)
	e2.name = "Kevin"
	fmt.Println("New ->", e2)
	// constructor
	e3 := NewEmployee(12, "kevin_", false)
	fmt.Println("Constructor ->", e3)
}
