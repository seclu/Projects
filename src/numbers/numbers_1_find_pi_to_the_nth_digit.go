package numbers

import (
	"fmt"
)

func FindPiToTheNthDigit() {
	var n int

	fmt.Println("Compute π to how many decimal places? (≤200) ")
	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		fmt.Println(err)
		return
	}

	if n > 200 {
		fmt.Println("Error:", n, "> 200 - too many decimal places!")
		return
	}

	if n < 0 {
		fmt.Println("Error:", n, "< 0")
		return
	}

	fmt.Printf("%f", calcNthDigit(n))
}

/*
Use Bailey–Borwein–Plouffe formula
https://en.wikipedia.org/wiki/Bailey%E2%80%93Borwein%E2%80%93Plouffe_formula
 */
func calcNthDigit(n int) float64 {
	var pi float64 = 0
	var pi16 float64 = 1

	for i := 0 ; i < n ; i++ {
		pi += float64(1)/float64(pi16) *
			(float64(4)/(float64(8)*float64(i) + float64(1)) -
				float64(2)/(float64(8)*float64(i) + float64(4)) -
				float64(1)/(float64(8)*float64(i) + float64(5)) -
				float64(1)/float64((8)*float64(i)+float64(6)))
		pi16 *= float64(16)
	}

	return pi
}