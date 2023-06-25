package main

import "fmt"

// create interface kosong
func random() interface{} {
	return 123131
}

func main() {
	fmt.Println("type assertions di golang")

	// result := random()
	// resultString := result.(string) // type assertion karena yakin hasilnya adalah string
	// fmt.Println(resultString)

	// resultInt := result.(int) // panic
	// fmt.Println(resultInt)

	// contoh type assertion atau pengubah tipe variabel tapi harus hati2
	// var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// type assertion dengan menggunakan switch agar lebih aman
	// ketika kita tidak tahu tipe datanya
	// tujuannya menghindari panic agar aplikasi tidak berhenti
	resultRandom := random()
	switch value := resultRandom.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	case bool:
		fmt.Println("bool", value)
	default:
		fmt.Println("unknown type")
	}

}
