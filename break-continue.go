package main

import "fmt"

func main() {
	// break di gunakan untuk menghentikan perulangan
	// continue menhentikan perulangan yg saat itu sedang berjalan

	fmt.Println("perulangna menggunakan break")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("perulangan ke ", i)
	}

	fmt.Println("Perulangan menggunakan continue")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("perulangan ke ", i)
	}
}
