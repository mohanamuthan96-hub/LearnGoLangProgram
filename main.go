package main

import (
	"fmt"
	"sync"
	"time"
)

//Branch - Slave0.1
//Buffered channels

func process(num int, wg *sync.WaitGroup) {
	fmt.Println("Added in the wait group ", num)
	time.Sleep(2 * time.Second)
	wg.Done()
	fmt.Println("Removed in the wait group ", num)
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
