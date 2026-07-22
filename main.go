package main

import (
	"fmt"
)

type Person struct {
	name string
}

type Worker interface {
	describe()
}

func (Per Person) describe() {
	fmt.Println("Name ", Per.name)
}

//Type Assertion
func describe(types any) {
	switch v := types.(type) {
	case int:
		fmt.Printf("I am a int and my value is %d\n", types.(int))
	case string:
		fmt.Printf("I am a string and my value is %s\n", types.(string))
	case Worker:
		v.describe()
	default:
		fmt.Printf("Unknowntype")
	}
}

func main() {
	describe("Naveen")
	describe(100)
	describe(Person{name: "Aravind"})
}
