package main

import "fmt"

type Permanent struct {
	name   string
	id     int
	salary int
	pf     int
}

type Contract struct {
	name   string
	id     int
	salary int
}

type Freelancer struct {
	name     string
	id       int
	noofhour int
	noofdays int
}

type Salarycalculator interface {
	Calculatesalary() int
}

func (per Permanent) Calculatesalary() int {
	sum := per.salary + per.pf
	return sum
}

func (con Contract) Calculatesalary() int {
	sum := con.salary
	return sum
}

func (fre Freelancer) Calculatesalary() int {
	sum := fre.noofdays + fre.noofhour
	return sum
}
func Calculatingsalary(salaries []Salarycalculator) int {
	sum := 0
	for _, value := range salaries {
		sum = sum + value.Calculatesalary()
	}
	return sum
}
func main() {
	per1 := Permanent{name: "Mohan", id: 1, salary: 10000, pf: 10000}
	per2 := Permanent{name: "Ram", id: 2, salary: 12000, pf: 10000}
	con1 := Contract{name: "Arun", id: 01, salary: 10000}
	free1 := Freelancer{name: "Kajul", id: 001, noofhour: 3, noofdays: 4}
	var employees []Salarycalculator
	employees = []Salarycalculator{per1, per2, con1, free1}
	result := Calculatingsalary(employees)
	fmt.Println(result)

}
