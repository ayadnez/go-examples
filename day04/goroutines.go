package main

import (
	"fmt"
	"time"
)

func greet() {
	time.Sleep(3 * time.Second)
	fmt.Println("hello")
}

func greet2() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello 2")
}

func main() {

	greet()

	go greet()
	go greet2()
	var input string

	fmt.Scanln(&input)

	fmt.Println("you have entered : ", input)

	fmt.Println("done")
}
