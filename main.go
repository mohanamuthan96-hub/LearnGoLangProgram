package main

import (
	"fmt"
	"math/rand/v2"
)

//Branch - Slave0.1

func Square(sq chan int, number int) {
	digit := make(chan int)
	go Numbergeneration(number, digit)
	sum := 0
	for value := range digit {
		sum = sum + value*value
	}
	sq <- sum
}

func Cube(cb chan int, number int) {
	digit := make(chan int)
	go Numbergeneration(number, digit)
	sum := 0
	for value := range digit {
		sum = sum + value*value*value
	}
	cb <- sum
}

func Numbergeneration(number int, digit chan int) {
	for number != 0 {
		value := number % 10
		digit <- value
		number = number / 10
	}
	close(digit)
}

var square chan int
var cube chan int

func main() {
	threedigit := rand.IntN(999)
	fmt.Println("Generate number ", threedigit)
	square = make(chan int)
	cube = make(chan int)
	go Square(square, threedigit)
	go Cube(cube, threedigit)
	fmt.Println(<-square, <-cube)

}
