package main

import (
	"fmt"
)

// 创建结构体
type Person struct {
	Id   int
	Name string
}

func (p Person) LogName() {
	fmt.Println(p.Name)
}

func (p *Person) ModifyName(newName string) {
	p.Name = newName
}

func main() {
	var tom Person = Person{Id: 0, Name: "Tom"}

	tom.LogName()

	tom.ModifyName("Mother")

	tom.LogName()
}
