package main

import (
	"fmt"
)


func sayHello()  {
	fmt.Println("Hello")
}
func anyHello(grerting string)  {
	fmt.Println("Hello")
	fmt.Println(grerting)
}

func cal(x int, y int)  {
	fmt.Println(x + y)
}
func cal2(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("test")
	sayHello()
	anyHello("good morning")
	cal(45, 89)
	cal_result := cal2(4, 6)
	fmt.Println((cal_result))
}