// Bad pattern

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float32
}

func (c Circle) area() {
	//出力と計算をしている SRPの違反
	fmt.Printf("circle area: %f\n", math.Pi * c.radius * c.radius)
}

type Square struct {
	sideLen float32
}

func (s Square) area() {
	//出力と計算をしている SRPの違反
	fmt.Printf("square area: %f\n", s.sideLen * s.sideLen)
}

func main() {
	c := Circle{radius: 5}
	c.area()

	s := Square{sideLen: 6}
	s.area()
}