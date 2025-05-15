package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("The highest factor is:", findTheHighestCommonFactor([]int{100, 36, 28, 12}))
}

func findTheHighestCommonFactor(numbers []int) int {
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	factors := []int{}
	for i := 1; i*i <= min; i++ {
		if min%i == 0 {
			factors = append(factors, i)
			if i != min/i {
				factors = append(factors, min/i)
			}
		}
	}

	sort.Ints(factors)
	fmt.Println("Factors of", min, "are:", factors)
	return len(factors)
}
