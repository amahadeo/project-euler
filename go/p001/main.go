package main

import "fmt"

// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6
// and 9. The sum of these multiples is 23. Find the sum of all the multiples of 3 or 5
// below 1000.

var limit = 1000
var multiples = []int{3, 5}

func main() {
	var sum int
	for i := 1; i < limit; i++ {
		for _, m := range multiples {
			if multipleOf(i, m) {
				sum += i
				break
			}
		}
	}
	fmt.Printf("Sum of multiples %v under %v: %v", multiples, limit, sum)
}

func multipleOf(num int, mult int) bool {
	if mult > num {
		return false
	}

	return num%mult == 0
}
