package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
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

	var arg = os.Args[1]

	var num int
	var err error

	if num, err = strconv.Atoi(arg); err != nil {
		panic(err)
	}

	var i int64 = 3
	a := 2

	for a < num {
		i += 2
		if isPrime(i) {
			a++
		}
	}

	fmt.Println(i)

}
