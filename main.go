package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("The highest factor of 100 is:", findFactors(100))
}

func findFactors(number int) int {
	factors := []int{}
	for i := 1; i*i <= number; i++ {
		if number%i == 0 {
			factors = append(factors, i)
			if i != number/i {
				factors = append(factors, number/i)
			}
		}
	}

	sort.Ints(factors)
	fmt.Println("Factors of", number, "are:", factors)
	return len(factors)
}
