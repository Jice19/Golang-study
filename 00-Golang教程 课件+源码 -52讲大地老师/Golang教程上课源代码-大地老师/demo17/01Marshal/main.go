/*
大地老师Golang专栏：https://www.itying.com/category-90-b0.html
Golang全家桶实战教程：https://www.itying.com/goods-1201.html
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	Name   string //私有属性不能被json包访问
	Sno    string
}

func main() {

	var s1 = Student{
		ID:     12,
		Gender: "男",
		Name:   "李四",
		Sno:    "s0001",
	}
	fmt.Printf("%#v\n", s1) //main.Student{ID:12, Gender:"男", Name:"李四", Sno:"s0001"}

	jsonByte, _ := json.Marshal(s1)
	jsonStr := string(jsonByte)
	fmt.Printf("%v", jsonStr) //{"ID":12,"Gender":"男","Name":"李四","Sno":"s0001"}
}
