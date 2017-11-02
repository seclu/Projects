package main

import (
	"./numbers"
	"./data_structures"
	"os"
	"fmt"
	"io/ioutil"
)

func main()  {
	var program, path, query string
	if len(os.Args) > 1 {
		program = os.Args[1]
		path = os.Args[2]
		query = os.Args[3]
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
		case "5":
			numbers.NextPrimeNumber()
		case "6":
			data, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			s := data_structures.Search{Text: string(data)}
			result := s.Search(query)
			fmt.Println(result)
		default:
			fmt.Println("Choose one program.")
	}
}
