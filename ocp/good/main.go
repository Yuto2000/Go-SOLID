// Good pattern

package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float32
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	slideLen float32
}

func (s square) area() float32 {
	return s.slideLen * s.slideLen
}

type triangle struct {
	height float32
	base float32
}

func (t triangle) area() float32 {
	return t.base * t.height / 2
}

type calculator struct {
	total float32
}

func (c calculator) sumAreas (shapes ...shape) float32 {
	var sum float32

	for _, value := range shapes {
		sum += value.area()
	}
	return sum
}

func main() {
	c := circle{radius: 5}
	s := square{slideLen: 6}
	t := triangle{height: 3, base: 4}
	calc := calculator{}
	fmt.Println("total of areas:", calc.sumAreas(c, s, t))
}