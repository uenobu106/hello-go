package fuga

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	// dx, dy => 256, 256

	lists := make([][]uint8, dy)

	fmt.Println(dx, dy)

	/*for y := 0; y < dy; y++ {
		list[y] = make([]uint8, dx)
	}
	*/
	
	for y := range lists {
		lists[y] = make([]uint8, dx)
	}

	// for y := 0; y < dy; y++ {
	// 	for x := 0; x < dx; x++ {
	// 		// list[y][x] = uint8(x + y)
	// 		// list[y][x] = uint8(x - y)
	// 		// list[y][x] = uint8(x * y)
	// 		// list[y][x] = uint8(x ^ y)
	// 		list[y][x] = uint8((x+y)/2)
	// 	}
	// }
	// return list


	for y, s := range lists {
		for x := range s {
			s[x] = uint8((x+y)/2)
		}
	}
	return lists
}

func main() {
	pic.Show(Pic)
}
