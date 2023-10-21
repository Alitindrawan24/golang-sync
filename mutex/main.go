package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	Mutex sync.Mutex
	Value int
}

func main() {
	counter := Counter{
		Mutex: sync.Mutex{},
		Value: 0,
	}
	group := &sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				defer group.Done()
				group.Add(1)

				counter.Mutex.Lock()
				counter.Value = counter.Value + 1
				counter.Mutex.Unlock()
			}
		}()
	}

	group.Wait()
	fmt.Println("Counter : ", counter.Value)
}
