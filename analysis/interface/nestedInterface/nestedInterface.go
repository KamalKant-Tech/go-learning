package main

import "fmt"

type Salaried interface {
	getSalary() int
}

type salary struct {
	basic, insurance, allowance int
}

func (s *salary) getSalary() int {
	return s.basic + s.insurance + s.allowance
}

type employee struct {
	firstName, lastname string
	salary              Salaried
}

func main() {
	emp := employee{
		firstName: "kamal",
		lastname:  "Kant",
		salary:    &salary{1100, 500, 300},
	}

	fmt.Println(emp.salary.getSalary())
}
