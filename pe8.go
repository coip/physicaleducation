package main

import (
	"fmt"
	"os"
	"strconv"
)

func prod(chunk string) int {

	i := 1
	var product int
	var tmp int
	var err error

	if product, err = strconv.Atoi(string(chunk[0])); err != nil {
		panic(err)
	}

	for i < len(chunk) {
		if tmp, err = strconv.Atoi(string(chunk[i])); err != nil {
			panic(err)
		}
		product *= tmp
		i++
	}

	return product

}

func main() {

	var arg = os.Args[1]

	var size int
	var err error

	var cmax int

	var hefty string = os.Args[2]

	if size, err = strconv.Atoi(arg); err != nil {
		panic(err)
	}

	fmt.Println(size)

	//if cmax, err = strconv.ParseInt(hefty[0:size], 10, 64); err == nil {
	//	fmt.Println(cmax)
	//} impl for just arbitrary-many adjacent making largest val

	cmax = prod(hefty[0:size])

	offset := 0
	var tmp int

	for offset+size < len(hefty) {
		tmp = prod(hefty[offset : offset+size])
		if tmp > cmax {
			cmax = tmp
		}
		offset++
	}
	fmt.Println(cmax)
}
