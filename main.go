package main

import (
	"fmt"
)

//Type Assertion
func describeint(types any) {
	value, ok := types.(int)
	if ok {
		fmt.Printf("The type is %T and value is %v and value of assertion is %v\n", types, types, value)
	}
	fmt.Printf("The type is %T and value is %v and value of assertion is %v\n", types, types, value)
}

func describestring(types any) {
	value := types.(string)
	fmt.Printf("The type is %T and value is %v and value of assertion is %v\n", types, types, value)
}

func describestruct(types any) {
	value := types.(any)
	fmt.Printf("The type is %T and value is %v and value of assertion is %v", types, types, value)
}
func main() {
	type Employee struct {
		name string
	}
	var sum interface{} = 100
	describeint(sum)
	var sums interface{} = "100"
	describeint(sums)
	var name any = "Mohan"
	describestring(name)
	var employee interface{} = Employee{name: "Selvi"}
	describestruct(employee)
}
