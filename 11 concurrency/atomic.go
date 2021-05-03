package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)


func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutine", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter,1)
			fmt.Println(counter)
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Goroutine:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutine:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}