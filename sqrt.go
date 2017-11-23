package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	last := 0.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-last) < .01 {
			break
		}
		last = z
	}
	return z
}

func main() {
	fmt.Printf("The square route of 25 is ~%v\n", Sqrt(25))
}
