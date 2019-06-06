package main

import (
	"fmt"
	"os"
	"strconv"
)

func sumofsq(a int) (sum int) {

	sum = 0

	for i := 1; i <= a; i++ {
		sum += (i * i)
	}

	return
}

func sqofsum(a int) (sum int) {

	sum = 0

	for i := 1; i <= a; i++ {
		sum += i
	}

	sum *= sum

	return
}

func main() {

	var arg string = os.Args[1]
	var num int
	var err error

	if num, err = strconv.Atoi(arg); err != nil {
		panic(err)
	}

	fmt.Println(num)

	a := sumofsq(num)
	b := sqofsum(num)
	c := b - a

	fmt.Printf("sumofSq: %v;\nsqofSums: %v;\n\nDiff: %v\n", a, b, c)

}
