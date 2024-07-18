package main

import "fmt"

func Fizzbuzz(cnt chan int, msg chan string) {
	for {
		i := <-cnt

		switch {
		case i%15 == 0:
			msg <- "fizbuzz"

		case i%5 == 0:
			msg <- "fizz"

		case i%3 == 0:
			msg <- "buzz"

		default:
			msg <- fmt.Sprintf("%d", i)

		}
	}
}

func main() {

	count := make(chan int)
	message := make(chan string)

	go Fizzbuzz(count, message)

	for i := 1; i < 101; i++ {
		count <- i
		fmt.Println(<-message)
	}

}
