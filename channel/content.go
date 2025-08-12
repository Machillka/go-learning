package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	fmt.Printf("Start Working\n")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("上下文发送退出信号, 退出, Err:", ctx.Err())
			return
		default:
			fmt.Printf("Working...\n")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)

	time.Sleep(3 * time.Second)
	cancel()

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Main 结束")
}
