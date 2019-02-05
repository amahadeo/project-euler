package main

import "fmt"

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

var divLimit = 20

func main() {
	smallest := divLimit

	for {
		if divisibleUpTo(smallest, divLimit) {
			break
		}
		smallest++
	}

	fmt.Printf("The smallest number that is evenly divisible from 1 to %v is %v\n", divLimit, smallest)
}

func divisibleUpTo(dividend int, limit int) bool {
	for i := 1; i <= limit; i++ {
		if dividend%i != 0 {
			return false
		}
	}

	return true
}
