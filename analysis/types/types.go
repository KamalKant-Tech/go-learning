package main

import (
	cfmt "basic/utils"
	"fmt"
	"reflect"

	"github.com/fatih/color"
)

func learnType(t reflect.Type) string {
	if t != nil {
		return "Has a type!"
	}
	return "<nil>"
}

func main() {
	age := 10
	name := "Ben"
	height := 123.5

	color.Cyan(cfmt.StrRepeat("-", "Go pointer example", 7))
	fmt.Printf("The type of name is: %T\n", name)
	fmt.Printf("The type of age is: %T\n", age)
	fmt.Printf("The type of height is: %T\n", height)

	fmt.Print(learnType(reflect.TypeOf(age)))
}
