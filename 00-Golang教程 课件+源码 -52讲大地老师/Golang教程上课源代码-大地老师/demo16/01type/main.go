/*
大地老师Golang专栏：https://www.itying.com/category-90-b0.html
Golang全家桶实战教程：https://www.itying.com/goods-1201.html
*/
package main

import "fmt"

//自定义类型
type myInt int

// type myFn func(int, int) int

//类型别名
type myFloat = float64

func main() {

	var a myInt = 10

	fmt.Printf("%v %T\n", a, a) //10 main.myInt

	var b myFloat = 12.3
	fmt.Printf("%v %T", b, b) //12.3 float64
}
