package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d start\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func SendData(ch chan string) {
	ch <- "first Data"
	ch <- "Second Data"
	ch <- "3rd Data"
}

func OutputData(ch chan string) {
	var data string
	for {
		data = <-ch
		if data == "" {
			break
		}
		fmt.Println(data)
	}
}

func NoBufChannelSender(ch chan int) {
	fmt.Println("Data Preparing")
	ch <- 86
	fmt.Println("Data already been sent")
}

func ChannelDataReceiver(ch chan int) {
	time.Sleep(time.Second) // 假设复杂操作使得线程停留 1s
	fmt.Println("准备接收")
	fmt.Println(<-ch)
	fmt.Println("接收结束")
}

func main() {
	ch := make(chan int)

	ch <- 1

	fmt.Println(<-ch)
	// go NoBufChannelSender(ch)
	// go ChannelDataReceiver(ch)

	time.Sleep(2 * time.Second)
}
