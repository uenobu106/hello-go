package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string){
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(math.Pi)
	fmt.Println(add(6, 4))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(19))
}
