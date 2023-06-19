package main

import "fmt"

func main() {
	name := "Fajar"

	if name == "fajar" {
		fmt.Println("halo fajar")
	} else if name == "siagian" {
		fmt.Println("selamat malam siagian")
	} else {
		fmt.Println("Hi, kamu fajar yang F besar")
	}

	// if dengan short statement
	if length := len(name); length > 5 {
		fmt.Println("karakter nama panjang")
	} else {
		fmt.Println("karakter cakep")
	}
}
