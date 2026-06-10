/*
大地老师Golang专栏：https://www.itying.com/category-90-b0.html
Golang全家桶实战教程：https://www.itying.com/goods-1201.html
*/
package main

import "fmt"

//一个结构体中可以嵌套包含另一个结构体或结构体指针

//父亲结构体
type Animal struct {
	Name string
}

func (a Animal) run() {
	fmt.Printf("%v 在运动\n", a.Name)
}

//子结构体
type Dog struct {
	Age     int
	*Animal //结构体嵌套 继承
}

func (d Dog) wang() {
	fmt.Printf("%v 在旺旺\n", d.Name)
}

func main() {
	var d = Dog{
		Age: 20,
		Animal: &Animal{
			Name: "阿奇",
		},
	}
	d.run()
	d.wang()
}
