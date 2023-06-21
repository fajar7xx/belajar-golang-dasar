package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("fajar")
	fmt.Println(person)

	// nil di gunakan untuk melakukan pengecekan apa data nya kosong
	cekData := NewMap("")
	if cekData == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(cekData)
	}

}
