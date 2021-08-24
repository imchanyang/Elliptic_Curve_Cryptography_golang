package main

import (
	"fmt"
)

func int_to_bin(num int64) string {
	var bin string

	bin = fmt.Sprintf("%b", num)

	return bin
}

func exp(a int , b int64, n int) int {
	c := 0
	f := 1
	binB := int_to_bin(b)
	k := len(binB)

	for i := 0; i < k; i++ {
		c = 2 * c
		f = (f * f) % n
		if binB[i] == '1' {
			c = c + 1
			f = (f * a) % n
		}
	}
	return f
}