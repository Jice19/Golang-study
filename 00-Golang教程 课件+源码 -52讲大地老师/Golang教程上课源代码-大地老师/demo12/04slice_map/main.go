/*
大地老师Golang专栏：https://www.itying.com/category-90-b0.html
Golang全家桶实战教程：https://www.itying.com/goods-1201.html
*/
package main

import "fmt"

func main() {
	//我们想在切片里面放一系列用户的信息,这时候我们就可以定义一个元素为map类型的切片
	// var userinfo = []string{"张三", "李四"}

	var userinfo = make([]map[string]string, 2, 2)

	// fmt.Println(userinfo[0]) //map[]   map不初始化的默认值nil

	// fmt.Println(userinfo[0] == nil) //true

	if userinfo[0] == nil {
		userinfo[0] = make(map[string]string)
		userinfo[0]["username"] = "张三"
		userinfo[0]["age"] = "20"
		userinfo[0]["height"] = "180cm"
	}

	if userinfo[1] == nil {
		userinfo[1] = make(map[string]string)
		userinfo[1]["username"] = "李四"
		userinfo[1]["age"] = "22"
		userinfo[1]["height"] = "170cm"
	}

	fmt.Println(userinfo)

}
