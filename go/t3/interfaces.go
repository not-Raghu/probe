package main

import (
	"math"
)

type GetHumanData interface {
	getName() string
	getAge() int
}

type Human struct {
	name string
	age  int
}

func (h Human) getName() string {
	return h.name
}

func (h Human) getAge() int {
	return h.age
}

// func main() {
// 	h := Human{"raghu", 1}
// 	fmt.Println(h.getName())
// 	fmt.Print(h.getAge())
// }

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return float64(math.Pi * c.radius * c.radius)
}

// func main() {
// 	s := circle{12}

// }
