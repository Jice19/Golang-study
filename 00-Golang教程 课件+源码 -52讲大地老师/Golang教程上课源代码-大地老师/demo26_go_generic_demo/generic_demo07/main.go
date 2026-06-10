package main

import "fmt"

//定义一个Usber接口
type Usber interface {
	Start()
	Stop()
}

//手机
type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println(p.Name, "开始工作")
}
func (p Phone) Stop() {
	fmt.Println("phone 停止")
}

//相机
type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机 开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机 停止工作")
}

//电脑的结构体
type Computer struct {
	Name string
}

// 电脑的Work方法要求必须传入Usb接口类型数据
func (c Computer) Work(usb Usber) {
	usb.Start()
	usb.Stop()
}
func main() {
	phone := Phone{
		Name: "小米手机",
	}
	computer := Computer{}
	//把手机插入电脑的Usb接口开始工作
	computer.Work(phone)

	camera := Camera{}
	//把相机插入电脑的Usb接口开始工作
	computer.Work(camera)
}
