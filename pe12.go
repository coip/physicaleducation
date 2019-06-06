package main

import (
	"fmt"
)

func triangle_num(idx, lastval int) int {
	return lastval + idx
}

func factors(x int) (y []int) {

	var i int

	for i = 1; len(y) < 2 && i <= x; i++ {
		if x%i == 0 {
			y = append(y, i)
		}
	}

	if y[len(y)-1] == x {
		return
	}

	for (y[len(y)-2]*y[len(y)-1] != x) && (y[len(y)-1]*y[len(y)-1] != x) {
		if x%i == 0 {
			y = append(y, i)
		}
		i++
	}
	return

}

func main() {

	var maxflen int = 0
	var b int = 0

	for a := 1; a <= 12000; a++ {
		lastb := b
		if a%10000 == 0 {
			fmt.Println(a)
		}
		b = triangle_num(a, lastb)
		c := factors(b)
		if len(c)*2-2 > maxflen {
			maxflen = len(c)*2 - 2
			fmt.Printf("%v:%v(%v) -> %v\n", a, b, len(c)*2-2, c)
		}
	}

	fmt.Println(maxflen)

}
