package main

import "fmt"

func main() {
	var n int

	fmt.Println("Generate the Fibonacci sequence to requested number:")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fib(n))
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n - 1) + fib(n - 2)
}