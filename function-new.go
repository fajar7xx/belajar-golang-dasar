package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Medan", "Sumut", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "deli serdang"

	*address2 = Address{"NAD", "ACEH", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	fmt.Println(address4)
}
