/*
大地老师Golang专栏：https://www.itying.com/category-90-b0.html
Golang全家桶实战教程：https://www.itying.com/goods-1201.html
*/
package main

import "fmt"

func main() {
	// 1、make创建map类型的数据

	// var userinfo = make(map[string]string)
	// userinfo["username"] = "张三"
	// userinfo["age"] = "20"
	// userinfo["sex"] = "男"
	// // fmt.Println(userinfo)
	// fmt.Println(userinfo["username"])

	//2、map 也支持在声明的时候填充元素
	// var userinfo = map[string]string{
	// 	"username": "张三",
	// 	"age":      "20",
	// 	"sex":      "男",
	// }
	// fmt.Println(userinfo)

	// 3、第三种创建map类型数据的方法

	userinfo := map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
	}
	fmt.Println(userinfo)
}
