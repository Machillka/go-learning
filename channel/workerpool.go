package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	Id   int
	Data string
}

func Worker(id int, tasks chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for t := range tasks {
		fmt.Printf("Worker %d processing task %d, and data %s\n", id, t.Id, t.Data)
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	const POOLSIZE = 5

	tasks := make(chan Job, 10)
	var wg sync.WaitGroup

	// 启动 "工人"
	for i := 1; i <= POOLSIZE; i++ {
		wg.Add(1)
		go Worker(i, tasks, &wg)
	}

	for i := 1; i <= 50; i++ {
		// 模拟写入工作
		tasks <- Job{
			Id:   i,
			Data: fmt.Sprintf("payload-%d", i),
		}
	}

	close(tasks)
    wg.Wait()
    fmt.Println("all tasks done")
}
