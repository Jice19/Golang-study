/*
大地老师Golang专栏：https://www.itying.com/category-90-b0.html
Golang全家桶实战教程：https://www.itying.com/goods-1201.html
*/
package main

import "fmt"

//单向管道
func main() {
	// 1、在默认情况下下，管道是双向
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 12
	m1 := <-ch1
	m2 := <-ch1
	fmt.Println(m1, m2) //10 12

	// 2、管道声明为只写
	ch2 := make(chan<- int, 2)
	ch2 <- 10
	ch2 <- 12
	// <-ch2   //receive from send-only type chan<- int

	// 3、管道声明为只读

	ch3 := make(<-chan int, 2)
	// ch3 <- 23

}
