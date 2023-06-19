package main

import "fmt"

// closure adalah kemampuan sebuah function berinteraksi dengan data2 di sekitarnya dalam scope yang sama
// harap gunakan fitur closure ini dengan bijak saat membuat aplikasi

func main() {
	counter := 0
	increment := func() {
		counter := 0 // ini hanya local variable agar tidak mengakses variable di luarnya
		fmt.Println("increment")
		counter++
		fmt.Println(counter)
	}

	increment()
	increment()
	fmt.Println(counter)
}
