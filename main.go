package main

import (
	"fmt"
	"sync"
	"time"
)

type Change struct {
	mux   sync.RWMutex
	count int
}

func (mu Change) Get() int {
	mu.mux.RLock()
	defer mu.mux.RUnlock()
	return mu.count
}

func (mu *Change) Set() {
	mu.mux.Lock()
	defer mu.mux.Unlock()
	(*mu).count++
}
func result(i int, wg *sync.WaitGroup) {
	fmt.Println("Goroutine ", i)
	defer wg.Done()
	fmt.Println("Goroutine ", i, "Executed")
}
func main() {
	wg := &sync.WaitGroup{}
	count := &Change{
		count: 0,
	}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				count.Set()
			}
			wg.Done()
		}()

	}
	wg.Wait()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("Loop of", i, " ", count.Get())
			time.Sleep(2 * time.Millisecond)
			wg.Done()
		}()

	}
	wg.Wait()
}
