package main

import (
	"math/rand"
	"time"
)

func rand_validpoint() Point {
	Point1 := Point{x: 2, y: 6}
	Point2 := Point{x: 4, y: 19}
	Point3 := Point{x: 13, y: 23}
	Point4 := Point{x: 16, y: 2}
	Point5 := Point{x: 19, y: 16}
	Point6 := Point{x: 27, y: 2}
	Point7 := Point{x: 0, y: 7}
	Point8 := Point{x: 2, y: 23}
	Point9 := Point{x: 5, y: 7}
	Point10 := Point{x: 8, y: 19}
	Point11 := Point{x: 14, y: 6}
	Point12 := Point{x: 16, y: 27}
	Point13 := Point{x: 20, y: 3}
	Point14 := Point{x: 27, y: 27}
	Point15 := Point{x: 8, y: 10}
	Point16 := Point{x: 0, y: 22}
	Point17 := Point{x: 3, y: 1}
	Point18 := Point{x: 5, y: 22}
	Point19 := Point{x: 10, y: 4}
	Point20 := Point{x: 14, y: 23}
	Point21 := Point{x: 17, y: 10}
	Point22 := Point{x: 20, y: 26}
	Point23 := Point{x: 1, y: 5}
	Point24 := Point{x: 3, y: 28}
	Point25 := Point{x: 6, y: 12}
	Point26 := Point{x: 10, y: 25}
	Point27 := Point{x: 15, y: 2}
	Point28 := Point{x: 17, y: 19}
	Point29 := Point{x: 24, y: 7}
	Point30 := Point{x: 1, y: 24}
	Point31 := Point{x: 4, y: 10}
	Point32 := Point{x: 6, y: 17}
	Point33 := Point{x: 13, y: 6}
	Point34 := Point{x: 15, y: 27}
	Point35 := Point{x: 19, y: 13}
	Point36 := Point{x: 24, y: 22}

	valid_point := []Point {
		Point1, Point7, Point13,Point19, Point25, Point31,
		Point2, Point8, Point14,Point20, Point26, Point32,
		Point3, Point9, Point15,Point21, Point27, Point33,
		Point4, Point10, Point16,Point22, Point28, Point34,
		Point5, Point11, Point17,Point23, Point29, Point35,
		Point6, Point12, Point18,Point24, Point30, Point36,
	}

	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(35)

	return valid_point[a]
}

func validpoint(i int) Point {
	Point1 := Point{x: 2, y: 6}
	Point2 := Point{x: 4, y: 19}
	Point3 := Point{x: 13, y: 23}
	Point4 := Point{x: 16, y: 2}
	Point5 := Point{x: 19, y: 16}
	Point6 := Point{x: 27, y: 2}
	Point7 := Point{x: 0, y: 7}
	Point8 := Point{x: 2, y: 23}
	Point9 := Point{x: 5, y: 7}
	Point10 := Point{x: 8, y: 19}
	Point11 := Point{x: 14, y: 6}
	Point12 := Point{x: 16, y: 27}
	Point13 := Point{x: 20, y: 3}
	Point14 := Point{x: 27, y: 27}
	Point15 := Point{x: 8, y: 10}
	Point16 := Point{x: 0, y: 22}
	Point17 := Point{x: 3, y: 1}
	Point18 := Point{x: 5, y: 22}
	Point19 := Point{x: 10, y: 4}
	Point20 := Point{x: 14, y: 23}
	Point21 := Point{x: 17, y: 10}
	Point22 := Point{x: 20, y: 26}
	Point23 := Point{x: 1, y: 5}
	Point24 := Point{x: 3, y: 28}
	Point25 := Point{x: 6, y: 12}
	Point26 := Point{x: 10, y: 25}
	Point27 := Point{x: 15, y: 2}
	Point28 := Point{x: 17, y: 19}
	Point29 := Point{x: 24, y: 7}
	Point30 := Point{x: 1, y: 24}
	Point31 := Point{x: 4, y: 10}
	Point32 := Point{x: 6, y: 17}
	Point33 := Point{x: 13, y: 6}
	Point34 := Point{x: 15, y: 27}
	Point35 := Point{x: 19, y: 13}
	Point36 := Point{x: 24, y: 22}

	valid_point := []Point {
		Point1, Point7, Point13,Point19, Point25, Point31,
		Point2, Point8, Point14,Point20, Point26, Point32,
		Point3, Point9, Point15,Point21, Point27, Point33,
		Point4, Point10, Point16,Point22, Point28, Point34,
		Point5, Point11, Point17,Point23, Point29, Point35,
		Point6, Point12, Point18,Point24, Point30, Point36,
	}

	return valid_point[i]
}



func invalidpoint(i int) Point {

	Point1 := Point{x: 3, y: 6}
	Point2 := Point{x: 5, y: 19}
	Point3 := Point{x: 13, y: 25}
	Point4 := Point{x: 16, y: 4}
	Point5 := Point{x: 19, y: 17}
	Point6 := Point{x: 22, y: 2}
	Point7 := Point{x: 1, y: 7}
	Point8 := Point{x: 2, y: 22}
	Point9 := Point{x: 5, y: 5}
	Point10 := Point{x: 8, y: 17}
	Point11 := Point{x: 14, y: 7}
	Point12 := Point{x: 16, y: 25}
	Point13 := Point{x: 20, y: 0}
	Point14 := Point{x: 7, y: 10}

	invalid_point := []Point{
		Point1, Point7, Point13,
		Point2, Point8, Point14,
		Point3, Point9, Point4,
		Point10, Point5, Point11,
		Point6, Point12,
	}
	return invalid_point[i]
}
