package main_lib_simple

import "fmt"

type Person struct {
	firstname, lastname string
	age                 uint
}

func (p *Person) String() string {
	return fmt.Sprintf("%v %v (%v)" , p.firstname, p.lastname, p.age)
}

func (p *Person) Age() uint {
	return p.age
}

func (p *Person) Firstname() string {
	return p.firstname
}

func (p *Person) Lastname() string {
	return p.lastname
}

func NewPerson(firstname string, lastname string, age uint) *Person {
	return &Person{firstname: firstname, lastname: lastname, age: age}
}

