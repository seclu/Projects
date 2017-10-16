package numbers

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func PrimeFactorization(args ...float64) string {
	var x float64
	if (len(args) == 1) {
		x = args[0]
	}

	if x == 0.0 {
		fmt.Println("Enter a number and find all Prime Factors")
		_, err := fmt.Scanf("%f", &x)
		if err != nil {
			fmt.Println(err)
		}
	}

	result := "Calculated prime factors for a number " + strconv.FormatFloat(x, 'f', 0, 64) + " is/are: " + strings.Join(castElementsToString(calculatePrimeFactors(x)), ", ")
	fmt.Println(result)

	return result
}

func calculatePrimeFactors(x float64) []float64 {
	var e float64
	var primeFactors []float64

	var i float64 = 2
	e = math.Floor(math.Sqrt(x))

	for i <= e {
		if math.Mod(x, i) == 0 {
			primeFactors = append(primeFactors, i)
			x = x / i
			e = math.Floor(math.Sqrt(x))
		} else {
			i++
		}
	}

	if x > 1 {
		primeFactors = append(primeFactors, x)
	}

	return primeFactors
}

func castElementsToString(x []float64) []string {
	var primeFactors []string
	for _, element := range x {
		primeFactors = append(primeFactors, strconv.FormatFloat(element, 'f', 0, 64))
	}

	return primeFactors
}