package main

import (
	"fmt"
	"runtime"
)

func main() {
	var mem runtime.MemStats
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
	for i := 0; i < 1000000; i++ {

		runtime.ReadMemStats(&mem)

		fmt.Printf("Total allocated memory: %d kb \n", mem.TotalAlloc/1024)
		fmt.Printf("Number of memory allocations: %d\n", mem.Mallocs)
	}

}
