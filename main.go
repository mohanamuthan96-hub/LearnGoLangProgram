package main

import (
	"fmt"
	"time"
)

//Branch - Slave0.1

func server1(ser chan string) {
	time.Sleep(1 * time.Second)
	ser <- "Server 1 running"

}

func server2(ser chan string) {
	time.Sleep(1 * time.Second)
	ser <- "Server 2 running"
}

var ser1 chan string
var ser2 chan string

func main() {

	ser1 = make(chan string)
	ser2 = make(chan string)
	go server1(ser1)
	go server2(ser2)
	for {

		select {
		case condition1 := <-ser1:
			fmt.Println(condition1)
			return
		case condition2 := <-ser2:
			fmt.Println(condition2)
			return
		default:
			fmt.Println("Default Connection")
		}
	}
}
