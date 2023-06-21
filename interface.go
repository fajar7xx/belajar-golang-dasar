package main

import "fmt"

// implementasi interface

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

// function sayHello yang memiliki contract hasName
func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var person Person
	person.Name = "Nadiah"

	SayHello(person)

	cat := Animal{
		Name: "tomi",
	}
	SayHello(cat)
}
