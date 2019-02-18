package main

import (
	"fmt"
	"sync"
)

// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.

var limit = 2000000

func main() {
	var sum int
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 2; i < limit; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if isPrime(i) {
				ch <- i
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		sum += v
	}

	fmt.Printf("The sum of all primes under %v is %v\n", limit, sum)
}

func isPrime(x int) bool {
	for i := 2; i <= x/i; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
