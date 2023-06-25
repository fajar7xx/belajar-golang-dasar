package main

import (
	"errors"
	"fmt"
)

func main() {
	// var contohError error = errors.New("Ups error")
	// fmt.Println(contohError.Error())

	// karena pada funcion pembagi terdapat return 2 tipe data
	// maka akan return 2 juga
	// dia akan cek apaka err itu nill atau tidak
	// kalau nill dia akan memproses
	// kalau ada error maka tangkap errornya
	// silahkan liat function pembagi
	hasil, err := Pembagi(100, 10)
	if err == nil {
		fmt.Println("Hasil: ", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}

/*
return int dan error
*/
func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh nol")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}
