package main

import (
	"fmt"
	"sync"
)

func result(i int, wg *sync.WaitGroup) {
	fmt.Println("Goroutine ", i)
	defer wg.Done()
	fmt.Println("Goroutine ", i, "Executed")
}
func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 25; i++ {
		wg.Add(1)
		go result(i, &wg)
	}
	wg.Wait()
	fmt.Println("Gorotine finished")
}
