package main

import "fmt"

// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
//
// a2 + b2 = c2
// For example, 32 + 42 = 9 + 16 = 25 = 52.
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

var sum = 1000

func main() {
	for a := 1; a <= sum; a++ {
		for b := a; b <= sum-a; b++ {
			for c := b; c <= sum-(a+b); c++ {
				if validPythTrip(a, b, c) {
					fmt.Printf("The valid pythagorean triplet whose sum equals %v is %v and their product is %v\n", sum, []int{a, b, c}, a*b*c)
					return
				}
			}
		}
	}
	fmt.Println("Could not find a valid pythagorean triplet w/ the given constraints")
}

func validPythTrip(a, b, c int) bool {
	return a+b+c == sum && (a*a)+(b*b) == (c*c)
}
