package main

import "fmt"

func main() {

	// const di bagai di berbagai source code sesuai dengan kebutuhan yang ada
	const firstName string = "fajar"
	const lastName = "siagian"
	const value = 1000

	// deklarasikan multiple constant
	const (
		alamat = "jalan makmur"
		kota   = "deli serdang"
	)

	fmt.Println(firstName)
	fmt.Println(alamat)
}
