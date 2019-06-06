package main

import "fmt"

var alpha, beta, omega int = 1, 2, 3

func next(a, b int) (o int) {
	o = a + b
	return
}

func iterate() {
	alpha = beta
	beta = omega
	omega = next(alpha, beta)
}

func main() {

	sum := 0
	for omega < 4000000 {
		if alpha%2 == 0 {
			sum += alpha
		}
		iterate()
	}
	if beta%2 == 0 {
		sum += beta
	}
	if omega%2 == 0 {
		sum += omega
	}

	fmt.Printf("sum: %v\n", sum)
}
