package main

import "fmt"

//Empty Interface,any alias of interface
func describe(types interface{}) {
	fmt.Printf("The type is %T and value is %v\n", types, types)
}
func main() {
	type Employee struct {
		name string
	}
	var sum int = 100
	describe(sum)
	var name string = "Mohan"
	describe(name)
	employee := Employee{
		name: "Arun",
	}
	describe(employee)
}
