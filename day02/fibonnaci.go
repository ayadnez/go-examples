package main

import "fmt"

//  fucntion to get nth number in fibonacci series
func FibNumber(n int) int {
	f := make([]int, n+2)

	f[0], f[1] = 0, 1

	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// function to get fibnacci series upto n numbers
func Fib(n int) chan int {

	c := make(chan int)

	go func() {
		a, b := 0, 1

		for i := 0; i < n; i++ {
			a, b = b, a+b

			c <- a
		}
		close(c)
	}()

	return c
}

func main() {

	for x := range Fib(10) {
		fmt.Println(x)
	}

	Nfib := FibNumber(10)
	fmt.Println("nth number of fibbonci series is : ", Nfib)
}
