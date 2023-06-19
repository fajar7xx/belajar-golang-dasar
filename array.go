package main

import "fmt"

func main() {
	var names [3]string // isi array max 3 dengan type data string
	names[0] = "fajar"
	names[1] = "setiawan"
	names[2] = "siagian"

	// buatkan array dengan langsung isi datanya
	// dan max datanya adalah 3
	var values = [3]int{
		90,
		80,
		95,
	}

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	fmt.Println(values)

	// untuk mendapatkan panjang array
	fmt.Println("panjang array values = ", len(values))

}
