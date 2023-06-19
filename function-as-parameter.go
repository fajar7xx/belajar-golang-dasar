package main

import "fmt"

// function tidak hanay bisa kita simpan di dalam variable sebagai value
// namun juga bisa kita gunakan sebagai parameter untuk function lain

// kadang jika function terlalu panjang, agak riber untuk menuliskannya di dalam parameter
// type declaration juga bisa di gunakan utnuk membuat alias function, sehingga
// akan mempermudah kita menggunakan function sebagai parameter

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello", nameFiltered)
// }

// refactor from up
// type declaration
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "....."
	} else {
		return name
	}
}

func main() {
	filter := spamFilter
	sayHelloWithFilter("anjing", filter)
}
