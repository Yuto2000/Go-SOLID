// Bad pattern

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Circle struct {
// 	radius float32
// }

// func (c Circle) area() {
// 	//出力と計算をしている SRPの違反
// 	fmt.Printf("circle area: %f\n", math.Pi * c.radius * c.radius)
// }

// type Square struct {
// 	sideLen float32
// }

// func (s Square) area() {
// 	//出力と計算をしている SRPの違反
// 	fmt.Printf("square area: %f\n", s.sideLen * s.sideLen)
// }

// func main() {
// 	c := Circle{radius: 5}
// 	c.area()

// 	s := Square{sideLen: 6}
// 	s.area()
// }

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
