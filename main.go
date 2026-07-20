package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	lock  sync.Mutex
)

func highPriority() {
	lock.Lock()
	time.Sleep(200 * time.Millisecond)
	defer lock.Unlock()
	count = +10
}

func lowPriority() {
	lock.Lock()
	defer lock.Unlock()
	count = +1
}
func main() {
	for range 100 {
		go highPriority()
		go lowPriority()
	}
	time.Sleep(300 * time.Millisecond)
	fmt.Println(count)
}
