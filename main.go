package main

import (
	"fmt"
)

func VariableTest() {

	// 可以显式声明
	var i, j, k int
	// 可以隐式推导
	var a, b, c = 1, 1.5, true

	// 简短变量声明
	idx := 0

	// 简短变量可以重复声明, 但是不能全都是已经声明过的变量（不然就是赋值语句了）
	idx, value := 0, "Item i"

	// 指针类型
	p := &idx
	fmt.Println(p)
	// 解引用 修改 idx
	*p = 1145

	// 指向一块匿名地址
	p = new(int)
	*p = 1145
	fmt.Println(*p)

	fmt.Println(i, j, k, a, b, c, idx, value)
}

func main() {
	var greeting string = "Hello world!"
	fmt.Println(greeting)

	VariableTest()
}