package main

import (
	"fmt"
	"math"
)

// 構造体
type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Pointer receivers
// ポインタレシーバを使う2つの理由があります。
// ひとつは、メソッドがレシーバが指す先の変数を変更するためです。
// ふたつに、メソッドの呼び出し毎に変数のコピーを避けるためです。
// 例えば、レシーバが大きな構造体である場合に効率的です。

func (v *Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
	fmt.Println(Vertex.Abs(v))
	fmt.Printf("%T\n", Abs)
	fmt.Printf("%T\n", v.Abs)
	fmt.Printf("%T\n", Vertex.Abs)

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs2())
	v.Scale2(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs2())
}