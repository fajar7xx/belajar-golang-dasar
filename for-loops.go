package main

import "fmt"

func main() {
	count := 1

	for count <= 10 {
		fmt.Println("perulangan ke ", count)
		count++
	}

	// for dengan statemnt
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("percobaan ke ", counter)
	}

	// for range untuk data array, slice, map
	slice := []string{
		"fajar",
		"setiawan",
		"siagian",
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, name := range slice {
		fmt.Println("index ", index, "= ", name)
	}

	// kalau pada for range tidak menggunakan index atau i cukup _ (underscore)
	for _, value := range slice {
		fmt.Println(value)
	}

	// for range pada map
	person := make(map[string]string)
	person["name"] = "fajar siagian"
	person["title"] = "lead engineer"

	for key, value := range person {
		fmt.Println(key, "= ", value)
	}

}
