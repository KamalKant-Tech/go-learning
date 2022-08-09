package main

// Importing fmt package for the sake of printing
import (
	"fmt"
	_ "math" // Suppresses compiler warnings related to fmt if it is not being used, and executes initialization functions if there are any
	"math/cmplx"

	h "basic/analysis/human"
	_ "basic/analysis/operation"
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

var (
	goIsFun bool       = true
	maxInt  uint64     = 1<<64 - 1
	complex complex128 = cmplx.Sqrt(-5 + 12i)
)

var elements = map[string]map[string]string{
	"H": map[string]string{
		"name":  "Hydrogen",
		"state": "gas",
	},
	"He": map[string]string{
		"name":  "Helium",
		"state": "gas",
	},
	"Li": map[string]string{
		"name":  "Lithium",
		"state": "solid",
	},
	"Be": map[string]string{
		"name":  "Beryllium",
		"state": "solid",
	},
	"B": map[string]string{
		"name":  "Boron",
		"state": "solid",
	},
	"C": map[string]string{
		"name":  "Carbon",
		"state": "solid",
	},
	"N": map[string]string{
		"name":  "Nitrogen",
		"state": "gas",
	},
	"O": map[string]string{
		"name":  "Oxygen",
		"state": "gas",
	},
	"F": map[string]string{
		"name":  "Fluorine",
		"state": "gas",
	},
	"Ne": map[string]string{
		"name":  "Neon",
		"state": "gas",
	},
}

var sliceEx = make([]int, 3, 9)

type Artist struct {
	Name, Genre string
	Songs       int
}

func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}

func routeF(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	location := "Dorne" // Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	fmt.Println(firstName+" "+lastName, location)

	me := &Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)

	const f = "%T(%v)\n"
	fmt.Printf(f, goIsFun, goIsFun)
	fmt.Printf(f, maxInt, maxInt)
	fmt.Printf(f, complex, complex)

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

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

}
