package main

import (
	"fmt"
	"reflect"
	"strings"

	cfmt "basic/utils"

	"github.com/fatih/color"
)

var (
	var1 int    = 1234
	var2 string = "string"
)

type User struct {
	name       string
	occupation string
}

func main() {

	color.Cyan(cfmt.StrRepeat("-", "Go pointer example", 7))
	x := &var1
	fmt.Println(x)  // this return the address(0x1149268) of values of var1 assined in memory
	fmt.Println(*x) // *x will return the value(1234) of address where var1 has stored

	argAsValue(var1)
	fmt.Println("Zeroval", var1)

	argAsReference(&var1)
	fmt.Println("Zeroval", var1)

	fmt.Println("Zeroval", &var1)

	testPointers()

	// Go pointer to struct
	u := User{"John Doe", "gardener"}
	fmt.Println(u)

	modify(&u)

	fmt.Println(u.name)

	color.Cyan(strRepeat("-", "Go pointer example End", 7))

	// Go pointer with new keyword
	// The new keyword takes a type as an argument, allocates enough memory to fit a value of that type and returns a pointer to it.
	color.Cyan(strRepeat("-", "Go pointer with new keyword start", 7))
	var pu *User = new(User)
	fmt.Println(pu)
	fmt.Println(reflect.TypeOf(pu))

	pu.name = "Robert Roes"
	pu.occupation = "accountant"
	fmt.Println(pu)

	color.Cyan(strRepeat("-", "Go pointer with new keyword End", 7))
	// Go pointer to pointer
	color.Cyan(strRepeat("-", "Go pointer to pointer Start", 7))
	var a = 7
	var p = &a
	var pp = &p
	var ppp = &pp

	fmt.Println(a)
	fmt.Println(&a)

	fmt.Println("--------------------")

	fmt.Println(p)
	fmt.Println(&p)

	fmt.Println("--------------------")

	fmt.Println(pp)
	fmt.Println(&pp)

	fmt.Println("--------------------")
	fmt.Println(ppp)
	fmt.Println(&ppp)

	fmt.Println("--------------------")

	fmt.Println(*pp)
	fmt.Println(**pp)
	fmt.Println(***ppp)
	color.Cyan(strRepeat("-", "Go pointer to pointer End", 7))
}

func modify(pu *User) {

	pu.name = "Robert Roe"
	pu.occupation = "driver"
}

func argAsValue(v int) {
	v = 0
}

func argAsReference(v *int) {
	*v = 0 // it will dereference the 0 as value in the reference of the passes arguments
}

func testPointers() {
	var x *int = &var1
	// x = "fsdklfs"
	fmt.Println(*x)

}

func strRepeat(text string, midText string, n int) string {
	return strings.Repeat("\n", 1) + strings.Repeat(text, n) + strings.Repeat(midText, 1) + strings.Repeat(text, n)
}
