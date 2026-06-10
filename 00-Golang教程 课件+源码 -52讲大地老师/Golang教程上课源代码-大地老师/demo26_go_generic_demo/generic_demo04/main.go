package main

import "fmt"

type Number interface {
	int | int64 | float64
}

// 定义泛型方法
func Add[T Number](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(Add(2, 2.5))

	// fmt.Println(Add("1","2"))  //报错了
}
