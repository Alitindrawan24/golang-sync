package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	Value int64
}

func main() {
	counter := Counter{
		Value: 0,
	}
	group := &sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				defer group.Done()
				group.Add(1)

				atomic.AddInt64(&counter.Value, 1)
			}
		}()
	}

	group.Wait()
	fmt.Println("Counter : ", counter.Value)
}
