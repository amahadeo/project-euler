package main

import "fmt"

//A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 ? 99.

//Find the largest palindrome made from the product of two 3-digit numbers.

func main() {
	var largest int
	factors := make([]int, 2, 2)

	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			product := i * j
			if isPalindrome(product) && product > largest {
				largest = product
				factors[0], factors[1] = i, j
			}
		}
	}

	fmt.Printf("The largest palindrome created by the product of 3-digit numbers is %v, by multiplying %v and %v\n", largest, factors[0], factors[1])
}

func isPalindrome(orig int) bool {
	var rev int
	n := orig

	for n != 0 {
		remainder := n % 10
		rev = (rev * 10) + remainder
		n /= 10
	}
	return orig == rev
}
