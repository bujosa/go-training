package main

import (
	"fmt"
	"runtime"
	"sync"
)

// uses wait group
var wg sync.WaitGroup

// concurrency
func main() {
	fmt.Println("ARCH", runtime.GOARCH)
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	bar()

	wg.Wait()
}

func foo() {
   for i := 0; i < 10; i++ {
	   fmt.Println("foo:",i)
   }
   wg.Done()
}

func bar() {
   for i := 0; i < 10; i++ {
	   fmt.Println("bar:",i)
   }
}