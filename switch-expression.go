package main

import "fmt"

func main() {
	name := "fajar"

	switch name {
	case "fajar":
		fmt.Println("Halo fajar")
	case "siagian":
		fmt.Println("hallo bang")
	default:
		fmt.Println("Halo siapa disana")
	}

	// swith juga mendukup short statement
	//  sama seperti if

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("name terlalu panjang")
	case false:
		fmt.Println("nama sesuai")
	}
}
