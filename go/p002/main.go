package main

import "fmt"

// Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

var limit = 4000000

func main() {
	fibs := []int{}
	var sum, i int
	for {
		if i == 0 {
			fibs = append(fibs, 1)
			i++
			continue
		}

		if i == 1 {
			fibs = append(fibs, fibs[i-1])
			i++
			continue
		}

		new := fibs[i-1] + fibs[i-2]

		if new > limit {
			break
		}

		fibs = append(fibs, new)
		if new%2 == 0 {
			sum += new
		}

		i++

		fmt.Printf("Current Fibonacci Slice: %v, Sum of Evens: %v\n", fibs, sum)
	}

	fmt.Printf("This sum of all even fibonacci numbers up to %v is %v\n", limit, sum)
}
