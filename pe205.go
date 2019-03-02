package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	cdicecount = 6
	pdicecount = 9
	cdicesize  = 6
	pdicesize  = 4
)

func makeTimestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func init() {
	rand.Seed(makeTimestamp())
}
func rollplayer(c, r int) int {
	var cpoints = 0
	//	fmt.Println("##   rolling", c, r)
	for croll := 1; croll <= c; croll++ {
		randomint := 1 + rand.Intn(r)
		//		fmt.Printf("rolled a %d\n", randomint)
		cpoints += randomint

	}
	//	fmt.Println("done rolling", c, r)
	return cpoints
}
func main() {
	outcome := rollplayer(cdicecount, cdicesize) - rollplayer(pdicecount, pdicesize)
	if outcome == 0 {
		fmt.Println("T")
	} else if outcome > 0 {
		fmt.Println("C")
	} else if outcome < 0 {
		fmt.Println("P")
	}
}
