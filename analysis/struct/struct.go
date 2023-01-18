package main

import (
	"basic/analysis/struct/structpackage"
	cfmt "basic/utils"
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

type p structpackage.StrPackage
type n structpackage.NestedStruct

type FullNameTypes func(string, string) string

type Address struct {
	Name    string
	city    string
	Pincode int
	// StrPackage   p // embedded struct
	StrPackage   *p
	NestedStruct n // nested struct embedded in Address struct
}

type Employee struct {
	firstName, lastName string
	age, salary         int
	Docs                struct { // anonymous struct type
		DrivingLicense string
		Empid          int
		ExpLetter      string
	}
	Fullname FullNameTypes
	getName  func()
}

func (e Employee) getEmpSalary() int { // method to get the salary of the emp. It slightly diffrent from the as it recieve the reciever parama before the name of the function
	return e.salary
}

// Reiever Employee as pointers
func (e *Employee) getEmpDocs() []string { // method to get the salary of the emp. It slightly diffrent from the as it recieve the reciever parama before the name of the function
	return []string{
		e.Docs.DrivingLicense,
		strconv.Itoa(e.Docs.Empid),
	}
}

func main() {

	/** embedded struct implementation Start **/
	color.Cyan(cfmt.StrRepeat("-", "embedded struct implementation", 25))
	// Declaring a variable of a `struct` type
	// All the struct fields are initialized with their zero value
	var a Address
	fmt.Println(a)
	fmt.Printf("%+v\n", a)
	// Declaring and initializing a struct using a struct literal
	a1 := Address{Name: "Akshay", city: "Dehradun", Pincode: 3623572, StrPackage: &p{Age: 14, Occuption: "Software engineer"}} // embedded struct implementation

	fmt.Println("Address1: ", a1, a1.StrPackage.Age)

	color.Cyan(cfmt.StrRepeat("-", "embedded struct implementation End", 25))

	/** * embedded struct implementation Start  **/

	a2 := Address{Name: "Akshay", city: "Dehradun", Pincode: 3623572, NestedStruct: n{Designation: "Software engineer", S: structpackage.StrPackage{Age: 12, Occuption: "sfdsf"}}} // Naming fields while initializing a struct
	// a2.NestedStruct.S.Age = 25
	// a2.NestedStruct.S.Occuption = "Software engineer"

	fmt.Println("Address2: ", a2)

	// Uninitialized fields are set to their corresponding zero-value
	a3 := Address{Name: "Delhi"}
	fmt.Println("Address3: ", a3)

	color.Cyan(cfmt.StrRepeat("-", "Pointers to a struct", 7))
	// passing the address of struct variable
	// emp8 is a pointer to the Employee struct
	// Note: Your Docs is a field with anonymous struct type. As such, you have to repeat the type definition:
	emp8 := Employee{firstName: "Sam", lastName: "Anderson", age: 55, salary: 6000, Docs: struct {
		DrivingLicense string
		Empid          int
		ExpLetter      string
	}{DrivingLicense: "Yes", Empid: 123456, ExpLetter: "available"},
		Fullname: func(firstname string, lastname string) string {
			return firstname + " " + lastname
		},
	}
	// (emp8).firstName is the syntax to access
	// the firstName field of the emp8 struct
	fmt.Println("First Name:", emp8.firstName)
	fmt.Println("Emp Id:", emp8.Docs.Empid)
	fmt.Println("Age:", emp8.age)
	fmt.Println("Full Name:", emp8.Fullname(emp8.firstName, emp8.lastName))
	fmt.Println("Emp salary:", emp8.getEmpSalary())
	fmt.Println("Emp Docs:", emp8.getEmpDocs())

	color.Cyan(cfmt.StrRepeat("-", "Pointers to a struct end", 7))
}
