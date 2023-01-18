package main

// Importing fmt package for the sake of printing
import (
	"fmt"
	_ "math" // Suppresses compiler warnings related to fmt if it is not being used, and executes initialization functions if there are any

	h "basic/analysis/human"
	_ "basic/analysis/operation"
	jwtAuth "basic/analysis/operation/auth"
	v "basic/analysis/vehicle"

	_ "github.com/google/wire"
)

var (
	firstName           string = "Kamal"
	lastName                   = "Kant" // If an initializer is present, the type can be omitted, the variable will take the type of the initializer (inferred typing)
	occuption                  = "Software Engineer"
	name, location, age        = "Prince Oberyn", "Dorne", 32 // You can also initialize variables on the same line:
	action                     = func() {                     // A variable can contain any type, including functions:

	}
)

// constants
const (
	StatusOK                   = 200
	StatusCreated              = 201
	StatusAccepted             = 202
	StatusNonAuthoritativeInfo = 203
	StatusNoContent            = 204
	StatusResetContent         = 205
	StatusPartialContent       = 206
)

var sliceEx = make([]int, 3, 9)

func routeF(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	location := "Dorne" // Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	fmt.Println(firstName+" "+lastName, location)

	fmt.Println(len(sliceEx))

	// Goroutines
	go routeF(0)

	// Closure example
	x := 0
	increment := func() int {
		x++
		return x
	}

	fmt.Println("Increment", increment())

	var bmw v.Vehicle
	bmw = v.Car("World Top Brand")

	var labour h.Human
	labour = h.Man("Software Developer")

	for i, j := range bmw.Structure() {
		fmt.Printf("%-15s <=====> %15s\n", j, labour.Structure()[i])
	}
	jwtAuth.JwtAuthentication()
	jwtAuth.Authentication()

}
