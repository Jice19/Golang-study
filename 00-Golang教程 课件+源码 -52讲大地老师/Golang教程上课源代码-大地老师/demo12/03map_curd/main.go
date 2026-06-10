/*
大地老师Golang专栏：https://www.itying.com/category-90-b0.html
Golang全家桶实战教程：https://www.itying.com/goods-1201.html
*/
package main

import "fmt"

func main() {

	//map类型数据的curd
	//1、创建 修改map类型的数据
	// var userinfo = make(map[string]string)
	// userinfo["username"] = "张三"
	// userinfo["age"] = "20"
	// fmt.Println(userinfo)

	//2、创建 修改map类型的数据

	// var userinfo = map[string]string{
	// 	"username": "张三",
	// 	"age":      "20",
	// }

	// userinfo["username"] = "李四"
	// fmt.Println(userinfo)

	//3、获取 查找map类型的数据

	// var userinfo = map[string]string{
	// 	"username": "张三",
	// 	"age":      "20",
	// }
	// fmt.Println(userinfo["username"])  //获取

	// v, ok := userinfo["age"]
	// fmt.Println(v, ok)  //20 true

	// v, ok := userinfo["xxxxx"]
	// fmt.Println(v, ok) // 空 和 false

	//4、删除map数据里面的key以及对于的值

	var userinfo = map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
		"height":   "180cm",
	}
	fmt.Println(userinfo)

	delete(userinfo, "sex")
	fmt.Println(userinfo)

}
