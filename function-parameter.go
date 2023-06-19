package main

import "fmt"

// function parameter bisa lebih dari satu
// parameter juga tidak wajib
// namun jika kita menambahkan parameter pada function
// maka ketika memanggil function tsb wajib memasukkan parameternya

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hallo", firstName, lastName)
}

func main() {
	sayHelloTo("Fajar", "Siagian")
}
