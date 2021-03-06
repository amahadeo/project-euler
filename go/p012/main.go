package main

import (
	"fmt"
	"sync"
)

// The sequence of triangle numbers is generated by adding the natural numbers. So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28. The first ten terms would be:
//
// 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
//
// Let us list the factors of the first seven triangle numbers:
//
//  1: 1
//  3: 1,3
//  6: 1,2,3,6
// 10: 1,2,5,10
// 15: 1,3,5,15
// 21: 1,3,7,21
// 28: 1,2,4,7,14,28
// We can see that 28 is the first triangle number to have over five divisors.
//
// What is the value of the first triangle number to have over five hundred divisors?

type answer struct {
	n        int
	tri      int
	divCount int
}

var (
	divisors = 500
	lastTri  int
	batchNo  = 1000
)

func main() {
	var n int
	var ans answer

	for ans.tri == 0 {
		fmt.Printf("Checked %+v triangle numbers so far ...still checking\n", n)
		ans = batch(&n)
	}
	fmt.Printf("%v is the #%v triangle number and the first triangle number w/ more than %v divisors", ans.tri, ans.n, divisors)
}

func triangle(n int) int {
	lastTri += n
	return lastTri
}

func divCount(num int) int {
	var divCount int
	for i := 1; i <= num/i; i++ {
		if num%i == 0 {
			// Every time there is one evenly divisible divisor, there is a second - avoid double counting
			divCount += 2
		}
	}
	return divCount
}

func batch(n *int) answer {
	var wg sync.WaitGroup
	ansChan := make(chan answer)

	// batch checking of divisors for 1000 triangles at a time
	for i := 0; i < batchNo; i++ {
		*n++

		tri := triangle(*n)

		if tri < divisors {
			// A natural number can't be smaller than it's number of divisors
			continue
		}

		wg.Add(1)

		// lanch go routines to handle heavy math
		go func(n, tri int) {
			divs := divCount(tri)
			if divs > divisors {
				ansChan <- answer{n, tri, divs}
			}
			wg.Done()
		}(*n, tri)
	}

	go func() {
		wg.Wait()
		close(ansChan)
	}()

	var ans answer
	for a := range ansChan {
		// because concurrent go routines are used, sequence can be lost
		// it's necessary to compare answers if multiple solutions come through
		if ans.tri == 0 || ans.tri > a.tri {
			ans = a
		}
	}

	return ans
}
