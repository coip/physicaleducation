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

	for a := 1; a < 6000; a++ {
		b := triangle_num(a)
		c := factors(b)
		if len(c) > maxflen {
			maxflen = len(c)
			fmt.Printf("%v:%v(%v) -> %v\n", a, b, len(c), c)
		}
	}

	/*   for a := 1; maxflen <= 500; a++ {
	     b := triangle_num(a)
	     c := factors(b)
	     fmt.Printf("%v:%v(%v) -> %v\n", a, b, len(c), c)
	     if len(c) > maxflen { maxflen = len(c); fmt.Printf("%v:%v(%v) -> %v\n", a, b, len(c), c) }
	  }*/

}
