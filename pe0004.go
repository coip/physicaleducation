package main

import "fmt"

func reverse_int(n int) int {
	new_int := 0
	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}
	return new_int
}

func main() {

	a := 999
	b := 999

	for a > 500 {
		for b > 2 {
			if a*b == reverse_int(a*b) {
				fmt.Println(a * b)
			}
			b--
		}
		a--
		b = a
	}

}
