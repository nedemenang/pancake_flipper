package main

import (
	"os"
	"log"
	"fmt"
	"unicode/utf8"
)

func main() {

	var t, k int
	count := 0
	var s string

	stdin, err := os.Open("A-large-practice.in")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdin = stdin

	fmt.Scan(&t)

	for ; t > 0; t-- {
		flips := 0
		impossible := false;
		fmt.Scan(&s,&k)
	
		sCount := utf8.RuneCountInString(string(s))

		b := make([]bool, sCount)

	
		for i := 0; i < sCount; i++ {

			if string(s[i]) == "+" {
				b[i] = true
			} else {
				b[i] = false
			}
		}
		for i := 0; i < sCount; i++ {
			if !bool(b[i]) {
				if i+k <= sCount {
					flips++
					for j := i; j < i+k; j++ {
						b[j] = !b[j]
					}
				} else {
					impossible = true
					break
				}
			}
		}
		count++
		if impossible == false {
			fmt.Printf("Case #%v: %v\n", count, flips)
		} else {
			fmt.Printf("Case #%v: IMPOSSIBLE\n", count)
		}
	}
}
