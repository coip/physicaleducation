package main

import (
	"fmt"
	"math"
)

//pythagorean triplet : a^2 + B^2 == c^2. ie 3^2 + 4^2 = 5^2. a+b+c == 12. want a set that a+b+c == 1000. answer: product of a.b.c

func sum(a, b, c int) int {

	return a + b + c

}

func main() {

	var a, b, c int

	for a = 5; a < 500; a++ {
		for b = a + 1; b < 500; b++ {
			c = int(math.Sqrt(float64(a*a) + float64(b*b)))
			if (sum(a, b, c)) == 1000 {
				if a*a+b*b == c*c {
					fmt.Printf("a: %v | b: %v | c: %v\n", a, b, c)
				}
			}
		}
	}

}
