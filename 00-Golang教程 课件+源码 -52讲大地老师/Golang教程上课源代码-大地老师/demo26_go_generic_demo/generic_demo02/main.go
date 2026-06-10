package main

import "fmt"

func getDataInterface(value interface{}) interface{} {
	return value
}
func main() {
	str1 := getDataInterface("hello")
	fmt.Println(str1)
	//类型断言
	s, _ := str1.(string)
	//输出s的长度
	fmt.Println(len(s))

	num2 := getDataInterface(123)
	fmt.Println(num2)

	flag := getDataInterface(true)
	fmt.Println(flag)
}
