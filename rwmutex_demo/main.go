package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x       int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
	rwMutex sync.RWMutex
)

// writeWithLock 互斥锁的写操作
func writeWithLock() {
	mutex.Lock()
	x += 1
	time.Sleep(time.Millisecond * 10) // 假设写操作耗时10毫秒
	mutex.Unlock()
	wg.Done()
}

// readWithLock 互斥锁的读操作
func readWithLock() {
	mutex.Lock()
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	mutex.Unlock()
	wg.Done()
}

// writeWithRWLock 读写互斥锁的写操作
func writeWithRWLock() {
	rwMutex.Lock()
	x += 1
	time.Sleep(time.Millisecond * 10) // 假设写操作耗时10毫秒
	rwMutex.Unlock()
	wg.Done()
}

// readWithRWLock 读写互斥锁的读操作
func readWithRWLock() {
	rwMutex.Lock()
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwMutex.Unlock()
	wg.Done()
}

func do(wf, rf func(), wc, rc int) {
	start := time.Now()

	// wc个并发写操作
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf()
	}

	// rc个并发读操作
	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf()
	}

	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("x:%v const:%v\n", x, cost)
}

func main() {
	// 使用互斥锁 10并发写 1000并发读
	do(writeWithLock, readWithLock, 10, 1000)

	// 使用读写互斥锁 10并发写 1000并发读
	do(writeWithRWLock, readWithRWLock, 10, 1000)
}
