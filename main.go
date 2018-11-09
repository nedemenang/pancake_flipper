package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var t string
	var k int
	var count int
	impossible := false
	fmt.Print("Enter pancake string:\n")
	fmt.Scan(&t)

	fmt.Print("Enter flipper size:\n")
	fmt.Scan(&k)

	tCount := utf8.RuneCountInString(t)

	b := make([]bool, tCount)
	for i := 0; i < tCount; i++ {

		if string(t[i]) == "+" {
			b[i] = true
		} else {
			b[i] = false
		}
	}
	for i := 0; i < tCount; i++ {
		if !bool(b[i]) {
			if i+k <= tCount {
				count++
				for j := i; j < i+k; j++ {
					b[j] = !b[j]
				}
			} else {
				impossible = true
				break
			}
		}
	}
	if impossible == false {
		fmt.Printf("Number of flips: %v\n", count)
	} else {
		fmt.Print("IMPOSSIBLE\n")
	}
}
