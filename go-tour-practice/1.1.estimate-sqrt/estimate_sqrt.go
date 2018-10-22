// https://tour.go-zh.org/flowcontrol/8

package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 3.0
	for ; (z*z - x) > 0.000001; z -= (z*z - x) / (2 * z) { // newton
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
