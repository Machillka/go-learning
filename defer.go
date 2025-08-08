package main

import (
	"fmt"
)

func DeferTest() {
	fmt.Println("Defer 执行")
}

func Worker() int {
	fmt.Println("Do something 01")
	defer DeferTest()
	fmt.Println("Do something 02")

	return -1
}

func main() {
	fmt.Println("函数执行完毕, 返回值为", Worker())
}
