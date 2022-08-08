// Bad pattern

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}

type square struct {
	slideLen float32
}

type calculator struct {
	total float32
}

func (c calculator) sumAreas (shapes ...interface{}) float32 {
	var sum float32

	for _, value := range shapes {
		switch value.(type) {
		case circle:
			r := value.(circle).radius
			sum += math.Pi * r * r
		case square:
			l := value.(square).slideLen
			sum += l * l
		}
	}
	return sum
}

func main() {
	c := circle{radius: 5}
	s := square{slideLen: 6}
	calc := calculator{}
	fmt.Println("total of areas:", calc.sumAreas(c, s))
}