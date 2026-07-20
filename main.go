package main

import (
	"fmt"
	"sync"
)

var mutex1 sync.Mutex
var mutex2 sync.Mutex

func resource1(name string) {
	mutex1.Lock()
	fmt.Println(name, "Acquired resorde 1")
	resource2("Resource2")
	mutex1.Unlock()
}

func resource2(name string) {
	mutex2.Lock()
	fmt.Println(name, "Acquired resorde 2")
	resource1("Resource1")
	mutex2.Unlock()
}

func main() {
	go resource1("Resource1")

	go resource2("Resource2")

	select {}
}
