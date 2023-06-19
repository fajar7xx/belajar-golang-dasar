package main

import "fmt"

// biasa saat kita memberi tahu bahawa sebuah function mengembalikan value
// maka kita hanya mendeklarasikan tipe data return value di function
// namun kita juga bisa membut variable secara langsung di tipe data
// return functionnya

func gettCompleteName() (firstName, midleName, lastName string) {
	firstName = "Fajar"
	midleName = "Setiawan"
	lastName = "Siagian"

	return firstName, midleName, lastName
}

func main() {
	a, b, c := gettCompleteName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
