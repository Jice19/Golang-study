package main

import "fmt"

// Usber 泛型接口定义了USB设备的基本操作
// T 代表设备的具体类型
type Usber[T any] interface {
	Start()       // 启动设备
	Stop()        // 停止设备
	GetDevice() T // 获取具体的设备实例
}

// Phone 手机结构体，表示手机设备
type Phone struct {
	Name string // 手机名称
}

// Start 实现Usber接口的Start方法
func (p Phone) Start() {
	fmt.Println(p.Name, "开始工作")
}

// Stop 实现Usber接口的Stop方法
func (p Phone) Stop() {
	fmt.Println(p.Name, "停止")
}

// GetDevice 实现Usber接口的GetDevice方法
func (p Phone) GetDevice() Phone {
	return p
}

// Camera 相机结构体，表示相机设备
type Camera struct {
	Brand string // 相机品牌
}

// Start 实现Usber接口的Start方法
func (c Camera) Start() {
	fmt.Println(c.Brand, "相机 开始工作")
}

// Stop 实现Usber接口的Stop方法
func (c Camera) Stop() {
	fmt.Println(c.Brand, "相机 停止工作")
}

// GetDevice 实现Usber接口的GetDevice方法
func (c Camera) GetDevice() Camera {
	return c
}

// Computer 电脑结构体 - 使用泛型 ， T 代表连接到电脑的设备类型
type Computer[T any] struct {
	Name string // 电脑名称
}

// Work 电脑的Work方法使用泛型接口操作USB设备
// 该方法接受任何实现了Usber[T]接口的设备实例
// usb: 要操作的USB设备实例
func (c Computer[T]) Work(usb Usber[T]) {
	// 启动设备
	usb.Start()
	// 停止设备
	usb.Stop()

	// 通过GetDevice方法获取具体的设备信息并打印
	device := usb.GetDevice()
	fmt.Printf("设备信息: %+v\n", device)
}
func main() {
	// 创建手机设备实例
	phone := Phone{
		Name: "小米手机",
	}

	// 创建相机设备实例
	camera := Camera{
		Brand: "佳能",
	}

	// 创建适用于手机的电脑实例并执行工作流程
	computer := Computer[Phone]{Name: "itying的电脑"}
	computer.Work(phone)

	// 创建适用于相机的电脑实例并执行工作流程
	// 注意：对于不同类型的设备，需要创建对应类型的电脑实例
	computer2 := Computer[Camera]{Name: "itying的电脑"}
	computer2.Work(camera)
}
