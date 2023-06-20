package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// function struct
func (customer Customer) sayHi() {
	fmt.Println("halo customer", customer.Name)
}

func (customer Customer) showAge() {
	fmt.Println("umur anda saat ini", customer.Age)
}

func main() {
	var customer Customer
	customer.Name = "new customer"
	customer.Address = "medan"
	customer.Age = 20

	// panggil struct method
	customer.sayHi()
	customer.showAge()
}
