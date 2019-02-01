package main

import "fmt"

// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143 ?

var num = 600851475143
var primes []int

func main() {
	primeFactorsOf(num)
	fmt.Printf("The prime factors of %v are %v with the highest prime being %v\n", num, primes, highest(primes))
}

func primeFactorsOf(num int) {
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			primeFactorsOf(i)
			primeFactorsOf(num / i)
			isPrime = false
			break
		}
	}
	if isPrime {
		primes = append(primes, num)
	}
}

func highest(xi []int) int {
	h := 0
	for _, v := range xi {
		if v > h {
			h = v
		}
	}
	return h
}
