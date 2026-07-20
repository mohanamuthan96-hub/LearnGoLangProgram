package main

import (
	"fmt"
	"time"
)

//Branch - Slave0.1
//Buffered channels

func main() {
	buffer := make(chan int, 2)
	go write(buffer)
	time.Sleep(1 * time.Second)
	for v := range buffer {
		time.Sleep(1 * time.Second)
		fmt.Println("Read from channel", v)

	}
}

func write(buff chan int) {
	for i := 1; i < 6; i++ {
		buff <- i
		fmt.Println("Successfully wrote to buffer channel ", i)
	}
	close(buff)
}
