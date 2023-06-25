package main

import "fmt"

type Man struct {
	Name string
}

// tambahkan pointer
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name

	// check value
	fmt.Println("di method", man.Name)
}

func main() {
	fmt.Println("pointer pada method")

	fajar := Man{"Fajar"}
	fajar.Married()
	// agar datanya bisa berubah menjadi mr fajar maka tambahkan pointer pada struct method
	fmt.Println(fajar.Name)
}
