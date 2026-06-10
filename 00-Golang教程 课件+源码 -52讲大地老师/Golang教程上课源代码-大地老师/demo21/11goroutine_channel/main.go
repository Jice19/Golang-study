/*
大地老师Golang专栏：https://www.itying.com/category-90-b0.html
Golang全家桶实战教程：https://www.itying.com/goods-1201.html
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// 1、创建channel
	var ch1 = make(chan int, 3)
	wg.Add(1)
	go func() {
		for i := 1; i <= 3; i++ {
			num := <-ch1
			fmt.Println(num)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(time.Second)
			ch1 <- i
		}
		wg.Done()
	}()

	wg.Wait()
}
