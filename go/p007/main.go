package main

import "fmt"

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

// What is the 10 001st prime number?

// n here stands for the nth prime number
var n = 10001

func main() {
	primeCount, incrementor := 0, 1

	for primeCount < n {
		incrementor++

		if isPrime(incrementor) {
			primeCount++
		}
	}

	fmt.Printf("The #%v prime number is %v\n", n, incrementor)
}

func isPrime(num int) bool {
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
