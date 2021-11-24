package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	a := 0.0
	for i := 0; i < 10; i++ {
		a = z
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
		if math.Abs(z - a) < 0.000001 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
