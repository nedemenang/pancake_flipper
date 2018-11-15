package loop

import (
	"strings"
	"strconv"
)

func Loop(s string, k int) (result string) {

	flips := 0
	impossible := false
	b := strings.Split(s, "")
	for i := 0; i < len(b); i++ {
		if b[i] == "-" {
			if i+k <= len(b) {
				flips++
				for j := i; j < i+k; j++ {
					if b[j] == "+" {
						b[j] = "-"
					} else {
						b[j] = "+"
					}
				}
			} else {
				impossible = true
				break
			}
		}
	}
	if impossible == true {
		result = "IMPOSSIBLE"
	} else {
		result = strconv.Itoa(flips)
	}
	return result
}
