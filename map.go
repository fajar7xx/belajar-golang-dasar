package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "fajar",
		"address": "medan",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["title"] = "lead enginer"
	fmt.Println(person["title"])

	var book map[string]string = make(map[string]string)
	book["title"] = "belajar mencintaimu"
	book["author"] = "fajar siagian"
	book["isbn"] = "1333144222"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "isbn")

	fmt.Println(book)
	fmt.Println(len(book))
}
