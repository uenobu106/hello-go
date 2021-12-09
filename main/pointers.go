package main

import (
	"fmt"
)

func pointers() {
	var i,j int
	i, j = 42, 2701

	var p *int

	p = &i         // point to i
	fmt.Println(p) // read i through the pointer -> 0xc0000ae008
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
	fmt.Println(i) // see the new value of j
    
}
