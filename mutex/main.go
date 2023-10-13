package main

import (
	"fmt"
	"sync"
)

func main() {
	x := 0
	var mutex sync.Mutex
	group := &sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				defer group.Done()
				group.Add(1)

				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	group.Wait()
	fmt.Println("Counter : ", x)
}
