package main

import (
	"fmt"
)

func extended_euclid(a int, b int) ( int, int){
	if a == b {
		return 0, a
	} else if b == 0 {
		return 0, a
	} else {
		x1 := 1
		y1 := 0
		r1 := a

		x2 := 0
		y2 := 1
		r2 := b

		var q int
		var rT int
		var xT int
		var yT int

		for r2 != 0{

			q = r1 / r2

			rT = r1 - q * r2
			xT = x1 - q * x2
			yT = y1 - q * y2

			x1, y1, r1 = x2, y2, r2
			x2, y2, r2 = xT, yT, rT


		}
		return y1, r1
	}
	//	s, old_s := 0, 1
	//	t, old_t := 1, 0
	//	r, old_r := b, a
	//
	//	for r != 0 {
	//		quotient := old_r / r
	//		old_r, r = r, old_r - quotient * r
	//		old_s, s = s, old_s - quotient * s
	//		old_t, t = t, old_t - quotient * t
	//	}
	//
	//
	//	return old_s, old_t, old_r
	//}
}





func m_inv(a int, n int) int {
	if a < 0 {
		a += n
	}
	y, r := extended_euclid(n, a%n)
	// declared but not used

	if r != 1 {
		fmt.Println(r)
		fmt.Println("No multiplicative inverse")
		return 1
	} else {
		return y % n
	}
}