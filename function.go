package main

import "fmt"

func sayHello() {
	fmt.Println("hello kamu yang disaana")
}

func main() {
	// sayHello()

	for i := 0; i < 10; i++ {
		sayHello()
	}
}
