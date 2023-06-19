package main

import "fmt"

// function yang mengembalikan data
// kita harus menuliskan tipe data pengembaliannya
// maka function tsb wajib mengembalikan data

func getHello(name string) string {
	return "Hello " + name
}

func main() {
	result := getHello("Fajar Siagian")

	fmt.Println(result)
}
