// Bad pattern

package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	volume() float64
}

type square struct {
	sideLen float64
}

func (s square) area() float64 {
	return math.Pow(s.sideLen, 2)
}

func (s square) volume() float64 {
	return 0
}

type cube struct {
	sideLen float64
}

func (c cube) area() float64 {
	return math.Pow(c.sideLen, 2)
}

func (c cube) volume() float64 {
	return math.Pow(c.sideLen, 3)
}

func areaSum(sharpes ...shape) float64 {
	var sum float64
	for _, value := range sharpes {
		sum += value.area()
	}
	return sum
}

func areaVolumeSum(sharpes ...shape) float64 {
	var sum float64
	for _, value := range sharpes {
		sum += value.volume()
	}
	return sum
}

func main() {
	s := square{sideLen: 2}
	c := cube{sideLen: 3}

	fmt.Println(areaSum(s, c))
	fmt.Println(areaVolumeSum(s, c))

}