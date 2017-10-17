package numbers

import (
	"fmt"
	"math"
)

func NextPrimeNumber(args ...float64) {
	var n float64
	var answer string

	fmt.Println("Find next prime number")

	for {
		fmt.Print("Should I calculate next prime number? [Y/n]")
		_, err := fmt.Scanf("%s", &answer)
		if err != nil && err.Error() != "unexpected newline" {
			fmt.Println(err)
			break
		}

		if answer == "n" {
			break
		}

		for {
			n++
			
			if isPrime(n) {
				fmt.Println(n)
				break
			}
		}
	}
}

func isPrime(number float64) bool {
	var n float64 = math.Abs(number)
	var i float64 = 2

	for i <= math.Sqrt(n) {
		if math.Mod(n, i) == 0 {
			return false
		}

		i++
	}

	return true
}