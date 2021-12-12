package main

import "fmt"



func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	var i interface{} = "hello"

	/*
	// キャストする
	t := i.(T)

	// okには真偽が入る
	t, ok := i.(T)
	*/
	
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	fmt.Println("==========")
	do(21)
	do("hello")
	do(true)
	fmt.Println("==========")

	f = i.(float64) // panic
	fmt.Println(f)


}