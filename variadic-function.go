package main

import "fmt"

// parameter yg berada di posisi terakhir, memiliki kemampuan dijadikan sebuah  var args
// var args artinya datanya bisa menerima lebih dari satu input, atau anggap saya semacam array.
// apa bedanya dengan parameter biasa dengan tipe data array ?
//  1. jika parameter tipe array , kita wajib membuat array terlebih dahulu sebelum mengirim ke function
// 	2. jika parameter menggunakan var args, kita bisa langsung mengirim datanya, jika lebih dari satu, cukup gunakan tanda koma

func sumAllVar(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sumAllVar(10, 12, 25, 100, 23)
	fmt.Println("total dari sumallvar adalah ", total)

	newSlice := []int{10, 10, 10, 10}
	totalSlice := sumAllVar(newSlice...) // menggunakan titik tiga agar bisa menggunakan siice pada variadic
	fmt.Println(totalSlice)
}
