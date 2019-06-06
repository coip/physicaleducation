package main

import "fmt"

func passes(n int) bool {
	for i := 2; i <= 20; i++ {
		if n%i != 0 {
			if i > 15 {
				fmt.Printf("%v fails %v\n", n, i)
			}
			return false
		}
	}
	return true
}

func main() {

	num := 2520
	passed := false

	for passed == false {
		if num%25000 == 0 {
			fmt.Printf("trying %v\n", num)
		}
		if passes(num) {
			fmt.Printf("winna winna chicken dinna: %v", num)
			passed = true
		}
		num++
	}

}
