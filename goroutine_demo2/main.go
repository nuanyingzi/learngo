package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(i int) {
			fmt.Println("Hello ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
