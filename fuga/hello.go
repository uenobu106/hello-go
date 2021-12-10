package fuga

import (
	"fmt"
	"runtime"
)

func Hello() string {
	defer fmt.Println("world")

	fmt.Println("hello")
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	return "hoge"
}