package main

import (
	"log"
	"math"
	"time"
)

func main() {
	for i := 100000000; i > 0; { // yep, it's and endless loop ;)
		sum1, elapsed1 := slow(i)
		log.Printf("The slow() function took %v for make a sum of all numbers from 1 to %d,"+
			" and the sum was %d", elapsed1, i, sum1)

		sum2, elapsed2 := fast(i)
		log.Printf("The fast() function took %v for make a sum of all numbers from 1 to %d,"+
			" and the sum was %d", elapsed2, i, sum2)
		i = i * 10
	}
}

func slow(n int) (int, time.Duration) {
	start := time.Now()
	sum := 0
	for i := 0; i <= n; i++ { // Loop through all the numbers one by one
		sum += i
	}
	elapsed := time.Since(start)
	return sum, elapsed
}

func fast(n int) (int, time.Duration) {
	start := time.Now()
	sum := int(math.Pow(float64(n), 2))/2 + n/2 // Gauss summary. No loop needed.
	elapsed := time.Since(start)
	return sum, elapsed
}
