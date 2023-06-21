package main

import "fmt"

// implementasi interface kosong
// jadi bisa memasukkan tipe data apa saja
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func main() {
	var data interface{} = Ups(4)
	fmt.Println(data)
}
