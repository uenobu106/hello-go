package fuga

import "fmt"

type MapVertex struct {
	Lat, Long float64
}

var m map[string] MapVertex

func Map() {
	m = make(map[string] MapVertex)
	m["Bell Labs"] = MapVertex {
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}