package main

import (
	"fmt"
	"os"
	"strconv"
)

func collatz(x int) (y int) {
	if x%2 == 0 {
		return x / 2
	}
	return 3*x + 1
}

func main() {

	var arg string = os.Args[1]
	var num int
	var err error
	var c int
	var cmax int
	var ctmp int = 0

	//var seq []int

	if num, err = strconv.Atoi(arg); err != nil {
		panic(err)
	}

	//seq = append(seq, num)

	//fmt.Printf("collatz(%v): %v", num, num)

	for a := 2; a < num; a++ {
		for c = collatz(a); c != 1; c = collatz(c) {
			//fmt.Printf(" -> %v", c)
			//seq = append(seq, c)
			ctmp++
		}
		if ctmp > cmax {
			cmax = ctmp
			fmt.Printf("len(collatz(%v)) == %v\n", a, cmax)
		}
		ctmp = 0
	}
	//seq = append(seq, 1)
	//fmt.Println(" -> 1")

	//fmt.Printf("len(seq) == %v\n", len(seq))
	//fmt.Println(seq)

}
