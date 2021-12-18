package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// s1 := strings.Fields(s)
	m := make(map[string]int)

	for _, word := range strings.Fields(s) {
		// m[word] = m[word] + 1
		m[word]++
	}

	return m
}

func main() {
	fmt.Printf("Fields are: %q", strings.Fields("I am learning Go!"))
	wc.Test(WordCount)
}
