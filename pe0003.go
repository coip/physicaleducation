package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	var number int
	var err error
	if number, err = strconv.Atoi(arg); err != nil {
		panic(err)
	}

	arg = os.Args[2]
	var divattmpt int
	if divattmpt, err = strconv.Atoi(arg); err != nil {
		panic(err)
	}

	fmt.Println("Trying to factor down to primes:")
	fmt.Println(number)

	fmt.Println("divattmpt:")
	fmt.Println(divattmpt)

	for divattmpt < number/2 {
		if number%divattmpt == 0 {
			fmt.Printf("%v, ", divattmpt)
		}
		divattmpt++
	}

}
