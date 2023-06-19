package main

// function tidak hanya dapat mengembalikan satu value tapi juga
// dapat mengembalikan multiple value

import "fmt"

func getFullName() (string, string) {
	return "fajar", "siagian"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
