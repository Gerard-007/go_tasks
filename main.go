package main

import (
	"fmt"
)

func main() {
	fmt.Println("The highest factor is:", findTheHighestCommonFactor([]int{10, 20, 30}))
}

func findTheHighestCommonFactor(numbers []int) int {
	hcf := numbers[0]
	for _, num := range numbers[1:] {
		hcf = gcd(hcf, num)
		if hcf == 1 {
			break
		}
	}
	return hcf
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
