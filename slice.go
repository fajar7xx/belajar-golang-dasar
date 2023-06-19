package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// pada saat array di rubah maka slice akan berubah
	// begitu juga slice berubah maka array akan berubah
	// karena berdasarkan reference

	months[5] = "Juni baru"
	fmt.Println(slice1)

	slice1[0] = "bulan lahir fajar"
	fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "fajar")
	fmt.Println(slice3)
	slice3[1] = "bukan desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)
}
