package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
	message := recover()

	if message != nil {
		fmt.Println("error dengan message: ", message)
	}

}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error")
	}

	fmt.Println("aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("halo selamat datang")
}
