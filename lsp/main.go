package main

import "fmt"

type transport interface {
	getName() string
}

type vehicle struct {
	name string
}

func (v vehicle) getName() string {
	return v.name
}

type car struct {
	vehicle
	wheel int
	engine int
}

type motorcycle struct {
	vehicle
	wheel int
}

type printer struct {}

func (p printer) printTransportName(t transport) {
	fmt.Println("name: ", t.getName())
}

func main() {
	v := vehicle{
		name: "toyota",
	}

	c := car{
		vehicle: vehicle{name: "car-name"},
		wheel: 4,
		engine: 1,
	}

	m := motorcycle{
		vehicle: vehicle{name: "motorcycle-name"},
		wheel: 2,
	}

	p := printer{}
	p.printTransportName(v)
	p.printTransportName(c)
	p.printTransportName(m)
}