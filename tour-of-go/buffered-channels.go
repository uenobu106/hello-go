package main

import "fmt"

func main() {
	ch := make(chan int, 2) // makeの引数に値を入れてバッファの長さを与える
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
