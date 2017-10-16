package numbers

import (
	"fmt"
	"math"
	"strings"
	"strconv"
)

func FindEToTheNthDigit() {
	var n int
	var decimalsCap int = 200

	fmt.Println("Compute e how many decimal places (≤", decimalsCap, ")")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println(err)
		return
	}

	if n > decimalsCap {
		fmt.Println("Error:", n, "> 200 - too many decimal places!")
		return
	}

	var format = []string{"%.", strconv.Itoa(n), "f"}
	fmt.Printf(strings.Join(format, ""), math.E)
}
