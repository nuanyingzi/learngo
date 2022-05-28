package main

import (
	"fmt"
	"time"
)

// work_pool

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, job)
		results <- job * 2
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("worker:%d end job:%d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启3个goroutine
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	// 发送5个任务
	for j := 0; j < 5; j++ {
		jobs <- j
	}
	close(jobs)
	for k := 0; k < 5; k++ {
		ret := <-results
		fmt.Println(ret)
	}
}
