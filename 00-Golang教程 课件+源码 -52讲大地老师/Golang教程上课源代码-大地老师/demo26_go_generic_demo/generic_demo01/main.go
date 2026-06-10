package main

import "fmt"

// 函数重复
func getStringData(value string) string {
	return value
}
func getIntData(value int) int {
	return value
}
func getBoolData(value bool) bool {
	return value
}

func main() {
	// 使用特定类型的函数
	str := getStringData("this is str")
	fmt.Println(len(str)) // 输出: this is str

	num := getIntData(123)
	fmt.Println(num)

	b := getBoolData(true)
	fmt.Println(b)
}
