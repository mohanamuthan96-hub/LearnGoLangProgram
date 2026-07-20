package main

import (
	"fmt"
	"sync"
)

var mutexA sync.Mutex
var mutexB sync.Mutex

func resource1(name string) {
	mutexA.Lock()
	defer mutexA.Unlock()
	fmt.Println(name, " has mutex A")
	fmt.Println(name, " trying to acquire mutex B")
	if !mutexB.TryLock() {
		fmt.Println("Cannot Aquire mutex B, Calling off")
		for range 100 {

		}
		resource1(name)
	} else {
		defer mutexB.Unlock()
		fmt.Println(name, " acquired mutex B")
		fmt.Println("working")
	}
}

func resource2(name string) {
	mutexB.Lock()
	defer mutexB.Unlock()
	fmt.Println(name, " has mutex B")
	fmt.Println(name, " trying to acquire mutex A")
	if !mutexA.TryLock() {
		fmt.Println("Cannot Aquire mutex A, Calling off")
		for range 100 {

		}
		resource2(name)
	} else {
		defer mutexA.Unlock()
		fmt.Println(name, " acquired mutex A")
		fmt.Println("working")
	}

}

func main() {
	go resource1("Routine A")

	go resource2("Routine B")
	fmt.Println("Waiting")
	select {}
}
