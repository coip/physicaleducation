package main

import (
	"fmt"
	"math"
)

func sqrt(a int64) int64 {
	return int64(math.Sqrt(float64(a)))
}

func isPrime(a int64) bool {
	if a <= 1 {
		return false
	}
	for b := sqrt(a); b >= 1; b-- {
		if b == 1 {
			return true
		}
		if a%b == 0 {
			return false
		}
	}
	return true
}

func main() {

	var sum int = 0

	for a := 2; a < 2000000; a++ {
		if isPrime(int64(a)) {
			sum += a
		}
	}

	fmt.Println(sum)

}
