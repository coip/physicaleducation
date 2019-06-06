package main

import (
	"fmt"
)

func triangle_num(x int) (y int) {
	for i := 1; i <= x; i++ {
		y += i
	}
	return
}

func factors(x int) (y []int) {

	for i := 1; i < x; i++ {
		if x%i == 0 {
			y = append(y, i)
		}
	}
	y = append(y, x)

	return

}

func main() {

	var maxflen int = 0

	for a := 1; a < 1200; a++ {
		b := triangle_num(a)
		c := factors(b)
		if len(c) > maxflen {
			maxflen = len(c)
			fmt.Printf("%v:%v(%v) -> %v\n", a, b, len(c), c)
		}
	}

}
