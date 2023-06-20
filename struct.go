package main

import "fmt"

// deklarasi data dengan struct
type CustomerPremium struct {
	Name      string
	Address   string
	Age       int
	IsMarried bool
}

func main() {
	var fajar CustomerPremium
	fajar.Name = "fajar siagian"
	fajar.Address = "deli serdang"
	fajar.Age = 27

	fmt.Println(fajar)
	fmt.Println(fajar.Name)

	// struct literals
	joko := CustomerPremium{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     20,
	}
	fmt.Println(joko)

	budi := CustomerPremium{"budi setiadi", "Indonesia", 32, false}
	fmt.Println(budi)
}
