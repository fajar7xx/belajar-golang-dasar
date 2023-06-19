package main

import "fmt"

// function adalah first class citizen
// function juga merupakan tipe data, dan bisa di simpan di dalam variable

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodbye := getGoodBye
	fmt.Println(goodbye("Fajar"))
}
