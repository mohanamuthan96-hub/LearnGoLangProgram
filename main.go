package main

import "fmt"

type Person struct {
	name  string
	place string
}

type Worker interface {
	work() string
}

func (per Person) work() (name string) {
	fmt.Printf("The conrete type of interface %T and value id %v", per, per)
	return per.name
}

func main() {
	per := Person{
		name: "Mohan", place: "Vellore",
	}
	var worker Worker = per
	worker.work()
}
