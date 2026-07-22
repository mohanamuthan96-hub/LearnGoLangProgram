package main

import "fmt"

//Type Assertion
func describe(types any) {
	switch types.(type) {
	case int:
		fmt.Printf("I am a int and my value is %d\n", types.(int))
	case string:
		fmt.Printf("I am a string and my value is %s\n", types.(string))
	default:
		fmt.Printf("Unknowntype")
	}
}

func main() {
	describe("Naveen")
	describe(100)
	describe(struct{ name string }{name: "Aravind"})
}
