package main

import (
	"os"
	"log"
	"fmt"
	. "github.com/nedemenang/pancake_flipper/loop"
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
		
		fmt.Scan(&s,&k)

		result := Loop(s, k)
		count++
		fmt.Printf("Case #%v: %v\n", count, result)
	}
}




