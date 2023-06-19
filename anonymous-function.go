package main

import "fmt"

// biasanya kita membuatkan function dengan memberikan nama pada function tsb
// namun kadang ada kalanya lebih mudah membuat function secara lebih mudahmembuat function
// secara langsugn di variable atau parameter tanpa harus membuat function terlebih dahulu
// hal tersebut dinamakan anonymous function

// type declaration
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	// anonymous function
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("fajar", blacklist)

	// atau bisa begini
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
