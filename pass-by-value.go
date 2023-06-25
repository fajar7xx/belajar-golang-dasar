package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		"Deli Serdang",
		"Sumatra Utara",
		"Indonesia",
	}

	address2 := address1

	address2.City = "Medan"

	fmt.Println(address1)
	fmt.Println(address2)
}
