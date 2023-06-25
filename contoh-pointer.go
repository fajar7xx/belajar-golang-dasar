package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{
		"Medan", "Sumut", "Indonesia",
	}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Deli Serdang"

	// pointer * memaksa address1 untuk pindah memorty ke address2
	*address2 = Address{"Aceh", "Aceh", "Indonesia"}

	// siapapun yang mengacu pada address 1 karena ada pointer address2 yang melakukan perubahan maka akan
	// maka yang mengikuti pada address1 juga akan berubah

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
}
