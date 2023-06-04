package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpFajar NoKTP = "1207263005960011"
	var marriedStatus Married = true

	fmt.Println(noKtpFajar)
	fmt.Println(marriedStatus)
}
