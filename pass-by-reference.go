package main

import "fmt"

type Alamat struct {
	Kota     string
	Provinsi string
	Negara   string
}

func main() {
	alamat1 := Alamat{
		"Deli serdang",
		"Sumut",
		"Indonesia",
	}

	// alamat 2 ke pointer ke alamat1
	alamat2 := &alamat1 // alamat dua ini pointer ke alamat1

	alamat2.Kota = "Medan"

	fmt.Println(alamat1)
	fmt.Println(alamat2)

}
