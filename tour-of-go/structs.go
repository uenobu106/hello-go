package fuga

import "fmt"

// struct (構造体)は、フィールド( field )の集まり
// structのフィールドは、ドット( . )を用いてアクセスします。
type Vertex struct {
	X int
	Y int
}

func Struct() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
