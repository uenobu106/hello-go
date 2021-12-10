package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}


// map はキーと値とを関連付けます(マップします)。
// マップのゼロ値は nil です。 nil マップはキーを持っておらず、またキーを追加することもできません。
// make 関数は指定された型のマップを初期化して、使用可能な状態で返します。
var m map[string]Vertex


// mapリテラルは、structリテラルに似ていますが、 キー ( key )が必要です。
var m2 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}


// Map literals continued
// もし、mapに渡すトップレベルの型が単純な型名である場合は、リテラルの要素から推定できますので、その型名を省略することができます。
var m3 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	// fmt.Println(m["Bell Labs"])
	// fmt.Println(m["hoge"])
	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)

	m4 := make(map[string]int)

	m4["Answer"] = 42
	fmt.Println("The value:", m4["Answer"])

	m4["Answer"] = 48
	fmt.Println("The value:", m4["Answer"])

	delete(m4, "Answer")
	fmt.Println("The value:", m4["Answer"])


	// もし、 m に key があれば、変数 ok は true となり、存在しなければ、 ok は false となります。

    // なお、mapに key が存在しない場合、 elem はmapの要素の型のゼロ値となります。

    // Note: もし elem や ok をまだ宣言していなければ、次のように := で短く宣言できます:
    // elem, ok := m[key]
	
	v, ok := m4["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	
}