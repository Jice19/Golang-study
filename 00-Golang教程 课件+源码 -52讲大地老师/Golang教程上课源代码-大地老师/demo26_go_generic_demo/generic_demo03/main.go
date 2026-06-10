package main

//定义一个getData的泛型方法
func getData[T any](value T) T {
	return value
}

func main() {
	str1 := getData[string]("this is str")
	println(str1)

	num := getData[int](123)
	println(num)

	str2 := getData("this is str")
	println(len(str2))

}
