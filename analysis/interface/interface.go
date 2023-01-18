package main

import (
	"fmt"
	"reflect"
)

type Shape interface {
	area() float32
}

type Rectangle struct {
	length, breadth float32
}

func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

// access method of the interface
func calculate(s Shape) {
	fmt.Println(reflect.TypeOf(s))
	fmt.Println("Area:", s.area())
}

func main() {
	rect := Rectangle{2, 4}
	calculate(rect)
}
