package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func twopower(x int) *big.Int {
	return new(big.Int).Lsh(big.NewInt(1), 100)
	//return math.Pow(float64(2), ddfloat64(x)))
	/*if x == 1 { return 2 }
	  return 2*twopower(x-1)*/
}

func sumofindividuals(x *big.Int) (sum int) {

	//   str := strconv.FormatFloat(x, 10)
	sum = 0
	fmt.Println(len(str))

	for i := 0; i < len(str); i++ {
		sum += (int(str[i]) - 48)
	}

	return

}

func main() {
	thou := twopower(1000)
	fmt.Println(thou)
	fmt.Println(sumofindividuals(thou))
}
