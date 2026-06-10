package main

import "fmt"

//	定义一个泛型结构体 求出传入数据的最小值 []string{"1", "2", "3"}    []int{8, 2, 3, 5, 1}
type OrderedType interface {
	int | float64 | string
}
type MinStruct[T OrderedType] struct {
	list []T
}

func (m *MinStruct[T]) AddList(list []T) {
	m.list = list
}
func (m MinStruct[T]) GetMin() T {
	min := m.list[0]
	for _, v := range m.list {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	m1 := MinStruct[int]{}
	m1.AddList([]int{8, 2, 3, 5, 1, 2})
	fmt.Println("最小值:", m1.GetMin())

	m2 := MinStruct[string]{}
	m2.AddList([]string{"8", "2", "3"})
	fmt.Println("最小值:", m2.GetMin())
}
