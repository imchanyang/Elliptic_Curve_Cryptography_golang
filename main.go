package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Point struct {
	x int
	y int
}

const p int = 29
const a int = 4
const b int = 20
const (
	INFx = -1
	INFy = -1
)



func inf(P Point, I Point) bool {
	if P == I {
		return true
	} else {
		return false
	}
}



func valid(P Point) bool {
	x, y := P.x, P.y
	return exp(y, 2, p) == ((exp(x, 3, p) + a*x + b)%p)
}



func neg(P Point) Point {
	x, y := P.x, P.y
	negP := Point{x: x, y:p-y}
	return negP
}



func add(P Point, Q Point) Point {
	x3 := -1
	y3 := -1
	INF := Point{x: INFx, y: INFy}



	x1, y1 := P.x, P.y
	x2, y2 := Q.x, Q.y

	if Q == INF {
		x3 = x1
		y3 = y1
	} else if P == INF {
		x3 = x2
		y3 = y2
	} else if Q == neg(P) {
		x3, y3 = INF.x, INF.y
	} else if P == Q {
		x3, y3 = dbl(P).x, dbl(P).y
	} else {
		x3 =(((y2-y1)*m_inv((x2-x1), p))*((y2-y1)*m_inv((x2-x1), p))-x1-x2) % p
		y3 = (((y2-y1)*m_inv((x2-x1), p))*(x1-x3)-y1) % p

		if x3 < 0 {
			x3 += p
		}
		if y3 < 0 {
			y3 += p
		}
	}

		result := Point{x: x3, y: y3}
		return result
}


func dbl(P Point) Point {
	x3 := -1
	y3 := -1
	x1, y1 := P.x, P.y
	INF := Point{x: INFx, y: INFy}

	if P == INF {
		x3, y3 = INFx, INFy
	} else {
		x3 = ((((3*x1*x1)+a)*m_inv(2*y1, p))*(((3*x1*x1)+a)*m_inv(2*y1, p)) - (2*x1)) % p
		y3 = ((((3*x1*x1)+a)*m_inv(2*y1, p))*(x1-x3) - y1) % p

		if x3 < 0 {
			x3 += p
		}
		if y3 < 0 {
			y3 += p
		}
	}

	result := Point{x: x3, y: y3}
	return result
}


func mul(k int, P Point) Point {
	return binMul(k, P)
}

func binMul(k int, P Point) Point {
	Q := Point{x: INFx, y: INFy}
	if k == 0 {
		Q = Point{x: INFx, y: INFy}
	} else if k == 37 {
		Q = Point{x: INFx, y: INFy}
	} else if k == 38 {
		Q = P
	} else {
		k_ := int64(k)
		K := int_to_bin(k_)
		for i := 0; i < len(K); i++ {
			Q = dbl(Q)
			if K[i] == '1' {
				Q = add(Q, P)
			}
		}
	}

	return Q
}



func main() {
	score := 0
	g := Point{x: 1, y: 5}
	INF := Point{x: INFx, y: INFy}
	// validity of points
	for i := 0; i < 36; i++ {
		if valid(validpoint(i)) != true {
			fmt.Println(validpoint(i), "is a valid point", score)
			os.Exit(3)
		} else {
			score += 1
		}
	}
	for i := 0; i < 14; i++ {
		if valid(invalidpoint(i)) == true {
			fmt.Println(invalidpoint(i), "is not a valid point", score)
			os.Exit(3)
		} else {
			score += 1
		}
	}


	// addition (1) : Boundaries
	// identity
	if add(g, INF) != g {
		fmt.Println("Addition-Identity checking is failed", score)
		os.Exit(3)
	}
	score += 1
	if add(INF, g) != g {
		fmt.Println("Addition-Identity checking is failed", score)
		os.Exit(3)
	}
	score += 1

	// negative
	if add(g, neg(g)) != INF {
		fmt.Println("Addition-Negate checking is failed", score)
		os.Exit(3)
	}
	score += 1

	// same input
	if add(g, g) != dbl(g) {
		fmt.Println("Addition-Same input checking is failed", score)
		os.Exit(3)
	}
	score += 1


	// addition (2) : Random pairs
	for i := 0; i < 10; i++ {
		s1 := rand_validpoint()
		s2 := rand_validpoint()
		r := add(s1, s2)
		if inf(r, INF) != true && valid(r) != true {
			fmt.Println(s1, "+", s2, "=", r, "is not on the curve", score)
			os.Exit(3)
		}
		score += 1
	}

	// doubling (1) Boundaries
	// inf
	if dbl(INF) != INF {
		fmt.Println("Addition-Same input checking is failed", score)
		os.Exit(3)
	}

	// doubling (2) Random points
	for i := 0; i < 10; i++ {
		s := rand_validpoint()
		r := dbl(s)
		if valid(r) != true {
			fmt.Println("2P", s, "=", r, "in not on the curve", score)
			os.Exit(3)
		}
		score += 1
	}

	// multiplication
	// boundaries
	// 0p
	if mul(0, g) != INF {
		fmt.Println("Multiplication-0P is not INF", score)
		os.Exit(3)
	}
	score += 1

	// 36P + P = INF
	if add(mul(36, g), g) != INF {
		fmt.Println("Multiplication-31P+P is not INF", score)
		os.Exit(3)
	}
	score += 1

	// 37P = 0
	if mul(37, g) != INF {
		fmt.Println("Multiplication-32P is not INF", score)
		os.Exit(3)
	}
	score += 1

	// 38P + 1P
	if mul(38, g) != g {
		fmt.Println("Multiplication-33P should be same with P", score)
		os.Exit(3)
	}
	score += 1
	// random k
	for i := 0; i < 22; i++ {
		rand.Seed(time.Now().UnixNano())
		k := rand.Intn(36)
		Q := mul(k, g)
		if inf(Q, INF) != true && valid(Q) != true {
			fmt.Println(k, "P=", Q, "is not on the curve", score)
			os.Exit(3)
		}

		score += 1
	}

fmt.Println("All pass. Current score is", score)

}

