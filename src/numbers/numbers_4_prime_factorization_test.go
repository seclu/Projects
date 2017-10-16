package numbers

import "testing"

func TestPrimeFactorization(t *testing.T) {
	expected := "Calculated prime factors for a number 100 is/are: 2, 2, 5, 5"
	actual := PrimeFactorization(100)

	if expected != actual {
		t.Error("Test failed, expected: '%s', got: '%s'", expected, actual)
	}
}
