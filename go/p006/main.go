package main

import "fmt"

// The sum of the squares of the first ten natural numbers is,

// 12 + 22 + ... + 102 = 385
// The square of the sum of the first ten natural numbers is,

// (1 + 2 + ... + 10)2 = 552 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

var limit = 100

func main() {
	var sum, sumSquares int

	for i := 1; i <= limit; i++ {
		sum += i
		sumSquares += (i * i)
	}

	squareSum := sum * sum

	fmt.Printf("For all natural numbers up to %v, the square of the sum is %v, the sum of Squares are %v, and the difference between the two is %v\n", limit, squareSum, sumSquares, squareSum-sumSquares)
}
