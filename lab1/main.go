package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var y int32
	x := rand.Int31n(1000)

	switch {
	case x >= 5 && x <= 25:
		y = 10 / (x*x - 4)
	default:
		y = 8 - x*x - x*x*x
	}

	fmt.Printf("x: %d\n", x)
	fmt.Printf("Result: %d\n", y)
}
