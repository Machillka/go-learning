package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			data := i
			ch1 <- data
			ch2 <- "Hello world"
		}
	}()

	time.Sleep(500 * time.Millisecond)

	var flag = false
	for {
		select {
		case data := <-ch1:
			fmt.Printf("Case 1, read data %d from ch1\n", data)
		case data := <-ch2:
			fmt.Printf("Case 2, read data %s from ch2\n", data)
		case <-time.After(time.Second):
			fmt.Printf("超时")
			flag = true
			// default:
			// 	fmt.Printf("No data receive\n")
			// 	flag = true
		}

		if flag {
			break
		}
	}

	fmt.Printf("End of main")
}
