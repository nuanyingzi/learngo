package main

import (
	"fmt"
	"sync"
)

// sync.Map 并发安全的map

var (
	wg sync.WaitGroup
)

var m = make(map[int]int)
var m2 = sync.Map{}

func get(key int) int {
	return m[key]
}

func set(key, value int) {
	m[key] = value
}

func main() {
	/*for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			set(i, i*2)
			fmt.Printf("key:%v value:%v\n", i, get(i))
			wg.Done()
		}(i)
	}
	wg.Wait()*/
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			m2.Store(i, i*2)
			value, _ := m2.Load(i)
			fmt.Printf("key:%v value:%v\n", i, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
