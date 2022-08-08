// Good pattern

package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float32
}

type outPrinter struct{}

func (o outPrinter) toText (s shape) string {
	return fmt.Sprintf("the area is: %f", s.area())
}

type Circle struct {
	radius float32
}

func (c Circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

type Square struct {
	sideLen float32
}

func (s Square) area() float32 {
	return s.sideLen * s.sideLen
}

func main() {
	c := Circle{radius: 5}
	s := Square{sideLen: 6}
	o := outPrinter{}

	fmt.Println(o.toText(c))
	fmt.Println(o.toText(s))
}
