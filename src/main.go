package main

import (
	"./numbers"
	"os"
	"fmt"
)

func main()  {
	var program string
	if len(os.Args) > 1 {
		program = os.Args[1]
	}

	switch program {
		case "1":
			numbers.FindPiToTheNthDigit()
		case "2":
			numbers.FindEToTheNthDigit()
		case "3":
			numbers.Fibonacci()
		case "4":
			numbers.PrimeFactorization()
		default:
			fmt.Println("Choose one program.")
	}
}
