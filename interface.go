package main

import "fmt"

// implementasi interface

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	var person Person
	person.Name = "Nadiah"

	SayHello(person)
}
