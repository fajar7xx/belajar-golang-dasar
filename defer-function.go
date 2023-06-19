package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging() //karena menggunakan defer brrti berjalan di background
	fmt.Println("Run application")
}

func main() {
	runApplication()
}
