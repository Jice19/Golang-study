# Golang 学习笔记

> 基于大地老师 Golang 教程（52讲）全部 demo 代码整理

---

## 一、基础入门（demo01 ~ demo04）

### 1.1 输出函数

```go
// Println: 自动换行，多个值中间有空格
fmt.Println("go", "python", "php") // go python php

// Print: 不自动换行，多个值中间无空格
fmt.Print("go", "python", "php")   // gopythonphp

// Printf: 格式化输出，不自动换行
fmt.Printf("a=%v b=%v c=%v\n", a, b, c)
```

| 格式化动词 | 含义 |
|-----------|------|
| `%v` | 原样输出 |
| `%T` | 输出类型 |
| `%d` | 十进制整数 |
| `%b` | 二进制 |
| `%o` | 八进制 |
| `%x` | 十六进制 |
| `%f` | 浮点数 |
| `%.2f` | 保留2位小数 |
| `%c` | 字符 |
| `%t` | bool |
| `%p` | 指针地址 |
| `%#v` | Go语法格式输出 |
| `%02d` | 宽度2，不足补0 |

### 1.2 注释

```go
// 单行注释  (快捷键: Ctrl+/)

/*
  多行注释
  多行注释
*/
```

### 1.3 变量

```go
// 方式一：声明后赋值
var username string
username = "张三"

// 方式二：声明并初始化
var username string = "张三"

// 方式三：类型推导
var username = "张三"

// 方式四：短变量声明（仅函数内可用）
username := "张三"

// 一次声明多个变量
var a1, a2 string
var (
    username string
    age      int
    sex      string
)

// 匿名变量（用 _ 忽略返回值）
var username, _ = getUserinfo()
```

> **注意**：Go 语言中变量定义后必须使用；同一作用域不支持重复声明；短变量 `:=` 不能用于全局变量；变量名区分大小写。

### 1.4 常量与 iota

```go
// 常量定义后不可修改
const pi = 3.14159

// 多个常量
const (
    A = "A"
    B = "B"
)

// 省略值时与上一行相同
const (
    n1 = 100
    n2        // 100
    n3        // 100
)

// iota 常量计数器（const 内行索引）
const (
    n1 = iota  // 0
    n2         // 1
    n3         // 2
)

// iota 中间插队
const (
    n1 = iota  // 0
    n2 = 100   // 100
    n3 = iota  // 2
    n4         // 3
)

// 多常量一行
const (
    n1, n2 = iota + 1, iota + 2  // 1, 2
    n3, n4                        // 2, 3
)
```

---

## 二、数据类型（demo06）

### 2.1 整型（int）

| 类型 | 范围 | 占用空间 |
|------|------|---------|
| `int8` | -128 ~ 127 | 1字节 |
| `int16` | -32768 ~ 32767 | 2字节 |
| `int32` | -2^31 ~ 2^31-1 | 4字节 |
| `int64` | -2^63 ~ 2^63-1 | 8字节 |
| `uint8` | 0 ~ 255 | 1字节 |

- 64位系统上 `int` 默认是 `int64`，占用8个字节
- 使用 `unsafe.Sizeof()` 查看占用空间
- 高位转低位可能溢出，需注意

### 2.2 浮点型（float）

```go
var a float32 = 3.12   // 4字节
var b float64 = 3.12   // 8字节（默认）
f1 := 3.1345656456     // 默认 float64

// 科学计数法
3.14e2   // 3.14 × 10²
3.14e-2  // 3.14 ÷ 10²
```

> **精度丢失问题**：`1129.6 * 100 = 112959.99999999999`，金融计算推荐使用 `decimal` 包。

### 2.3 布尔型（bool）

- 只有 `true` 和 `false` 两个值
- 默认值为 `false`
- **不能**将整型强制转换为布尔型
- **不能**参与数值运算

```go
var flag = true
if flag {  // ✅ 正确
    fmt.Println("true")
}
// if 1 { }  // ❌ 错误，不能用整数代替bool
```

### 2.4 字符串（string）

```go
// 定义
var str1 string = "你好golang"
str2 := "你好go"

// 转义符：\n 换行，\\ 反斜杠，\" 双引号

// 多行字符串（反引号）
str := `this is str
this is str
this is str111`

// 长度（字节数）
len("aaaa")   // 4
len("你好")   // 6（UTF-8 编码，一个汉字3字节）

// 拼接
str3 := str1 + str2
str3 := fmt.Sprintf("%v %v", str1, str2)

// strings 包常用方法
strings.Split("123-456-789", "-")      // ["123" "456" "789"]
strings.Join([]string{"a","b"}, "*")   // "a*b"
strings.Contains("this is str", "is")  // true
strings.HasPrefix("this is str", "this") // true
strings.HasSuffix("this is str", "str") // true
strings.Index("this is str", "is")     // 2（首次出现位置）
strings.LastIndex("this is str", "is") // 5（最后出现位置）
```

### 2.5 字符（byte/rune）

```go
var a = 'a'
fmt.Printf("%v %c %T", a, a, a)  // 97 a int32

// byte 遍历（按字节）
s := "你好 golang"
for i := 0; i < len(s); i++ {
    fmt.Printf("%v(%c) ", s[i], s[i])
}

// rune 遍历（按字符，推荐处理中文）
for _, v := range s {
    fmt.Printf("%v(%c) ", v, v)
}

// 修改字符串（需转换为 []byte 或 []rune）
s1 := "big"
byteStr := []byte(s1)
byteStr[0] = 'p'
fmt.Println(string(byteStr))  // "pig"

s2 := "你好golang"
runeStr := []rune(s2)
runeStr[0] = '大'
fmt.Println(string(runeStr))  // "大好golang"
```

### 2.6 类型转换

```go
// 数字互转
int16(a) + b          // 整型之间
float64(a) + b        // 浮点之间
a + float32(b)        // 整型与浮点

// 其他转字符串：方式一 fmt.Sprintf
fmt.Sprintf("%d", 20)           // int → string
fmt.Sprintf("%.2f", 12.456)     // float → string
fmt.Sprintf("%t", true)         // bool → string
fmt.Sprintf("%c", 'a')          // byte → string

// 其他转字符串：方式二 strconv
strconv.FormatInt(int64(i), 10)      // int → string
strconv.FormatFloat(float64(f), 'f', 4, 32)  // float → string

// 字符串转数字
strconv.ParseInt("123456", 10, 64)     // string → int
strconv.ParseFloat("123.456", 64)      // string → float
```

---

## 三、运算符（demo07）

### 3.1 算术运算符

| 运算符 | 说明 |
|--------|------|
| `+ - * / %` | 加减乘除取余 |

```go
// 整数除法去掉小数部分
10 / 3   // 3
10.0 / 3.0  // 3.333...

// 取余公式：余数 = 被除数 - (被除数/除数)*除数
-10 % 3  // -1
10 % -3  // 1
```

> **`++` 和 `--`**：只能独立使用，不能赋值给其他变量（`a = i++` 错误）；只有后置写法（`++a` 错误）。

### 3.2 关系运算符

`==` `!=` `>` `>=` `<` `<=`

### 3.3 逻辑运算符

`&&` `||` `!`

> **短路**：`&&` 前为 false 则后面不执行；`||` 前为 true 则后面不执行。

### 3.4 赋值运算符

`=` `+=` `-=` `*=` `/=` `%=`

### 3.5 位运算符

| 运算符 | 说明 |
|--------|------|
| `&` | 按位与（两位为1才为1） |
| `\|` | 按位或（有一位为1就为1） |
| `^` | 异或（两位不同为1） |
| `<<` | 左移n位 = 乘以2^n |
| `>>` | 右移n位 = 除以2^n |

---

## 四、流程控制（demo08）

### 4.1 if 判断

```go
// 基本写法
if score >= 90 {
    fmt.Println("A")
} else if score > 75 {
    fmt.Println("B")
} else {
    fmt.Println("C")
}

// 带初始化语句（age 为局部变量）
if age := 34; age > 20 {
    fmt.Println("成年人")
}
```

> **注意**：`{` 必须紧挨条件语句；if 的 `{}` 不能省略。

### 4.2 for 循环

Go 语言**没有 while**，全部用 for 实现：

```go
// 标准 for
for i := 1; i <= 10; i++ { }

// 省略初始语句
i := 1
for ; i <= 10; i++ { }

// 省略初始和结束（类似 while）
i := 1
for i <= 10 {
    i++
}

// 死循环
for {
    if 条件 { break }
}
```

### 4.3 for range 遍历

```go
// 遍历字符串
for k, v := range "你好golang" { }

// 遍历切片
for _, val := range arr { }

// 遍历数组
for k, v := range arr { }
```

### 4.4 switch case

```go
// 基本写法（Go 中 break 可写可不写）
switch extname {
case ".html":
    fmt.Println("text/html")
case ".css":
    fmt.Println("text/css")
default:
    fmt.Println("找不到")
}

// 多值逗号分隔
switch n {
case 1, 3, 5, 7, 9:
    fmt.Println("奇数")
}

// 无表达式（类似 if-else）
switch {
case age < 24:
    fmt.Println("好好学习")
case age >= 24 && age <= 60:
    fmt.Println("好好赚钱")
default:
    fmt.Println("注意身体")
}

// fallthrough：穿透到下一个 case
case age >= 24:
    fmt.Println("好好赚钱")
    fallthrough  // 继续执行下一个 case
case age > 60:
    fmt.Println("这句也会执行")
```

### 4.5 break / continue / goto

```go
// break：跳出循环，支持 label 跳出指定循环
lable1:
for i := 0; i < 2; i++ {
    for j := 0; j < 10; j++ {
        if j == 3 { break lable1 }
    }
}

// continue：跳过当前循环进入下一次，也支持 label
// goto：无条件跳转
if n > 24 {
    goto lable3
}
lable3:
    fmt.Println("跳转到这里")
```

---

## 五、数组（demo09）

```go
// 定义
var arr1 [3]int
var arr2 = [3]int{23, 34, 5}
arr3 := [3]string{"php", "nodejs", "golang"}

// 编译器推断长度
arr4 := [...]int{12, 23, 45, 67}

// 指定索引初始化
arr5 := [...]int{0: 1, 1: 10, 5: 50}  // [1 10 0 0 0 50]

// 遍历
for i := 0; i < len(arr); i++ { }
for k, v := range arr { }

// 二维数组
var arr = [3][2]string{
    {"北京", "上海"},
    {"广州", "深圳"},
    {"成都", "重庆"},
}
```

> **注意**：
> - 数组长度是类型的一部分，`[3]int` 和 `[4]int` 是不同类型
> - **数组是值类型**：赋值给新变量时会复制整个数组

---

## 六、切片（demo10）

切片是**引用类型**，底层指向一个数组。

### 6.1 定义切片

```go
// 直接定义
var sliceA []int
sliceB := []string{"php", "java"}

// 基于数组定义
a := [5]int{55, 56, 57, 58, 59}
b := a[:]     // 所有值
c := a[1:4]   // [56 57 58]
d := a[2:]    // [57 58 59]
e := a[:3]    // [55 56 57]

// make 创建
sliceA := make([]int, 4, 8)  // 长度4，容量8
```

### 6.2 长度与容量

- **长度（len）**：切片包含的元素个数
- **容量（cap）**：从第一个元素到底层数组末尾的元素个数

```go
s := []int{2, 3, 5, 7, 11, 13}  // len=6 cap=6
a := s[2:]                        // len=4 cap=4
b := s[1:3]                       // len=2 cap=5
```

### 6.3 append 扩容

```go
sliceA = append(sliceA, 12)
sliceA = append(sliceA, 12, 23, 35)

// 合并切片（注意 ...）
sliceA = append(sliceA, sliceB...)
```

> **扩容策略**：容量 < 1024 时翻倍；>= 1024 时增长 1/4。

### 6.4 copy 复制

```go
sliceB := make([]int, 4, 4)
copy(sliceB, sliceA)  // 深拷贝，修改 sliceB 不影响 sliceA
```

### 6.5 删除元素

```go
// 删除索引为 2 的元素
a = append(a[:2], a[3:]...)
```

### 6.6 值类型 vs 引用类型

| 值类型 | 引用类型 |
|--------|---------|
| 基本数据类型、数组、结构体 | 切片、map、channel |
| 修改副本不影响原值 | 修改副本影响原值 |

---

## 七、Map（demo12）

### 7.1 定义与初始化

```go
// make 创建
userinfo := make(map[string]string)

// 声明时填充
userinfo := map[string]string{
    "username": "张三",
    "age":      "20",
}
```

> **注意**：map 未初始化时为 nil，不能直接赋值（会 panic），需要用 make 初始化。

### 7.2 CRUD 操作

```go
// 增/改
userinfo["username"] = "张三"
userinfo["username"] = "李四"  // 修改

// 查
fmt.Println(userinfo["username"])

// 安全查询（判断 key 是否存在）
v, ok := userinfo["age"]
if ok { /* key 存在 */ }

// 删
delete(userinfo, "sex")
```

### 7.3 遍历

```go
for k, v := range userinfo {
    fmt.Printf("key:%v value:%v\n", k, v)
}
```

### 7.4 切片与 Map 的组合

```go
// 元素为 map 的切片
var userinfo = make([]map[string]string, 2, 2)
userinfo[0] = make(map[string]string)  // map 需要单独初始化
userinfo[0]["username"] = "张三"

// map 值为切片
var userinfo = make(map[string][]string)
userinfo["hobby"] = []string{"吃饭", "睡觉", "敲代码"}
```

### 7.5 Map 排序

Map 是无序的，需要排序时：
1. 把 key 放入切片
2. 对切片排序（sort.Ints）
3. 按排序后的 key 遍历 map

### 7.6 Map 类型

Map 也是**引用类型**，直接赋值会共享底层数据。

---

## 八、函数（demo13）

### 8.1 基本语法

```go
// 函数定义
func sumFn(x int, y int) int {
    return x + y
}

// 参数简写（同类型参数）
func subFn1(x, y int) int { }

// 可变参数
func sumFn1(x ...int) int {
    sum := 0
    for _, v := range x { sum += v }
    return sum
}

// 固定参数 + 可变参数（可变参数放最后）
func sumFn2(x int, y ...int) int { }
```

### 8.2 多返回值

```go
func calc(x, y int) (int, int) {
    return x + y, x - y
}

// 命名返回值
func calc1(x, y int) (sum int, sub int) {
    sum = x + y
    sub = x - y
    return  // 裸返回
}
```

### 8.3 函数类型与高阶函数

```go
// 定义函数类型
type calcType func(int, int) int

// 函数作为参数
func calc(x, y int, cb calcType) int {
    return cb(x, y)
}

// 函数作为返回值
func do(o string) calcType {
    switch o {
    case "+": return add
    case "-": return sub
    }
}

// 匿名函数作为参数
calc(3, 4, func(x, y int) int { return x * y })
```

### 8.4 匿名函数与闭包

```go
// 匿名自执行函数
func() { fmt.Println("test..") }()

// 赋给变量
var fn = func(x, y int) int { return x * y }

// 闭包：函数内嵌套函数，内部函数可访问外部变量
func adder() func(int) int {
    var i = 10  // 常驻内存、不污染全局
    return func(y int) int {
        i += y
        return i
    }
}
fn := adder()
fn(10)  // 20
fn(10)  // 30
```

### 8.5 defer 延迟执行

```go
// 逆序执行（先 defer 的后执行）
defer fmt.Println(1)
defer fmt.Println(2)
defer fmt.Println(3)
// 输出: 3 2 1

// defer 注册时参数就确定了
x := 1
defer fmt.Println(x)  // 输出 1（不是后面的值）
x = 10

// 命名返回值 vs 匿名返回值
func f2() int {        // 匿名返回值
    var a int
    defer func() { a++ }()
    return a           // 返回 0（defer 修改的是 a，不是返回值）
}
func f3() (a int) {    // 命名返回值
    defer func() { a++ }()
    return a           // 返回 1（defer 修改的是返回值 a）
}
```

### 8.6 panic / recover

```go
func fn() {
    defer func() {
        err := recover()  // 捕获 panic
        if err != nil {
            fmt.Println("error:", err)
        }
    }()
    panic("抛出一个异常")
}
```

> `recover` 只有在 defer 调用的函数中才有效。

### 8.7 作用域

- **全局变量**：函数外定义，整个程序运行期间有效
- **局部变量**：函数内定义，只能在该函数内使用
- **块作用域**：if/for 语句内定义的变量只在该块内有效

---

## 九、指针（demo15）

```go
var a = 10
var p = &a        // p 是指针变量，类型 *int
fmt.Println(*p)   // 10（解引用，取出指针指向的值）
*p = 30           // 修改指针指向的值
fmt.Println(a)    // 30

// 函数传参：值传递不会修改原值，指针传递会
func fn1(x int) { x = 10 }    // 不影响原值
func fn2(x *int) { *x = 40 }  // 影响原值
```

### new 和 make

| | new | make |
|------|-----|------|
| 作用 | 分配内存，返回指针 | 初始化 slice/map/channel |
| 返回值 | 指针类型 | 引用类型本身 |
| 适用 | 任意类型 | slice/map/channel |

```go
a := new(int)  // *int，值为 0
var userinfo = make(map[string]string)  // 已初始化的 map
```

> 引用类型（slice、map、channel、指针）必须初始化后才能使用，否则是 nil。

---

## 十、结构体（demo16）

### 10.1 自定义类型

```go
type myInt int           // 新类型
type myFloat = float64   // 类型别名
```

### 10.2 结构体定义与实例化

```go
type Person struct {
    Name string
    Age  int
    Sex  string
}

// 方式一
var p1 Person
p1.Name = "张三"

// 方式二：new（返回指针）
p2 := new(Person)
p2.Name = "李四"

// 方式三：& 取地址
p3 := &Person{Name: "赵四", Age: 23}

// 方式四：键值对
p4 := Person{Name: "哈哈", Age: 20}

// 方式五：按顺序（不推荐）
p5 := &Person{"张三", 20, "男"}
```

### 10.3 方法（接收者）

```go
// 值接收者（不修改原值）
func (p Person) PrintInfo() {
    fmt.Printf("姓名：%v 年龄:%v\n", p.Name, p.Age)
}

// 指针接收者（可修改原值）
func (p *Person) SetInfo(name string, age int) {
    p.Name = name
    p.Age = age
}
```

> 非本地类型（别的包的类型）不能定义方法。

### 10.4 结构体嵌套（继承）

```go
type Animal struct { Name string }
func (a Animal) run() { fmt.Printf("%v 在运动\n", a.Name) }

// Dog 嵌套 Animal → 继承
type Dog struct {
    Age int
    Animal           // 匿名嵌套
}

d := Dog{Age: 20, Animal: Animal{Name: "阿奇"}}
d.run()              // 可直接调用 Animal 的方法
d.Name = "旺财"       // 可访问匿名结构体的字段
```

### 10.5 匿名字段与字段冲突

```go
// 匿名字段（直接用类型名做字段名）
type Person struct {
    string
    int
}

// 字段冲突：有多个嵌套结构体拥有相同字段名时，需显式指定
u.Address.AddTime = "2020-05-1"
u.Email.AddTime = "2020-06-1"
```

### 10.6 结构体字段类型

```go
type Person struct {
    Name  string
    Hobby []string       // 切片字段
    map1  map[string]string  // map 字段
}
// 切片和 map 字段需要用 make 初始化后才能使用
p.Hobby = make([]string, 3, 6)
p.map1 = make(map[string]string)
```

### 10.7 值类型确认

**结构体是值类型**，赋值给新变量时会复制整个结构体。

---

## 十一、JSON 处理（demo17）

### 11.1 Marshal（序列化）

```go
type Student struct {
    ID     int
    Gender string
    Name   string
}

s1 := Student{ID: 12, Gender: "男", Name: "李四"}
jsonByte, _ := json.Marshal(s1)
jsonStr := string(jsonByte)
// {"ID":12,"Gender":"男","Name":"李四"}
```

### 11.2 Unmarshal（反序列化）

```go
var str = `{"ID":12,"Gender":"男","Name":"李四"}`
var s1 Student
err := json.Unmarshal([]byte(str), &s1)
```

### 11.3 结构体标签（Tag）

```go
type Student struct {
    ID     int    `json:"id"`
    Gender string `json:"gender"`
    Name   string `json:"name"`
}
// 序列化后 key 使用 tag 中的名字: {"id":12,"gender":"男","name":"李四"}
```

> 字段首字母**大写**才能被 json 包访问（公有属性）。

---

## 十二、包管理（demo18/demo19）

### 12.1 包的基本规则

```go
// 首字母大写 = 公有（可被外部访问），首字母小写 = 私有
package calc
var Age = 20           // 公有
var aaa = "私有变量"    // 私有
func Add(x, y int) int // 公有方法
```

### 12.2 init 函数

- 包导入时自动执行 `init()`
- 执行顺序：导入包 init → 当前包 init → main()

### 12.3 Go Modules

```bash
go mod init 项目名称       # 初始化项目
go mod tidy               # 下载当前缺少的依赖
```

### 12.4 导入第三方包

```go
import (
    "github.com/shopspring/decimal"  // 精确小数
    "github.com/tidwall/gjson"       // JSON 解析
)
```

---

## 十三、接口（demo20）

### 13.1 接口定义与实现

```go
type Usber interface {
    start()
    stop()
}

// Phone 实现 Usber 接口（隐式实现，无需声明）
func (p Phone) start() { fmt.Println(p.Name, "启动") }
func (p Phone) stop()  { fmt.Println(p.Name, "关机") }

// 接口作为变量类型
var p1 Usber = Phone{Name: "华为手机"}
p1.start()
```

### 13.2 接口作为函数参数（多态）

```go
func (c Computer) work(usb Usber) {
    usb.start()
    usb.stop()
}
computer.work(phone)
computer.work(camera)
```

### 13.3 空接口

```go
// 空接口可以表示任意类型
type A interface{}

var a A
a = "string"
a = 20
a = true

// 作为函数参数
func show(a interface{}) { }

// 作为 map 值
map[string]interface{}{}

// 作为切片元素
[]interface{}{1, 2, "你好", true}
```

### 13.4 类型断言

```go
// 方式一：ok 模式
v, ok := a.(string)
if ok { fmt.Println("string类型:", v) }

// 方式二：switch type（仅限 switch）
switch x.(type) {
case int:    fmt.Println("int类型")
case string: fmt.Println("string类型")
case bool:   fmt.Println("bool类型")
}
```

### 13.5 接口接收者

- **值接收者**：值类型和指针类型都可赋值给接口
- **指针接收者**：只能指针类型赋值给接口

### 13.6 接口嵌套

```go
type Animaler interface {
    Ainterface
    Binterface
}
```

---

## 十四、Goroutine 与 Channel（demo21）

### 14.1 Goroutine

```go
go test()  // 开启一个协程
```

> 主 goroutine 退出后，所有子 goroutine 都会退出。

### 14.2 Channel（管道）

```go
// 创建
ch := make(chan int, 3)  // 容量为3的有缓冲管道

// 存入
ch <- 10

// 取出
a := <-ch

// 管道是引用类型
```

> **阻塞**：管道满了再写会 deadlock；管道空了再读也会 deadlock。

### 14.3 单向管道

```go
// 只写管道
func write(ch chan<- int) { }

// 只读管道
func read(ch <-chan int) { }
```

### 14.4 select 多路复用

```go
select {
case v := <-intChan:
    // 从 intChan 读取
case v := <-stringChan:
    // 从 stringChan 读取
default:
    // 都没有数据时执行
}
```

### 14.5 并发安全

```go
// sync.Mutex（互斥锁）
var mutex sync.Mutex
mutex.Lock()
// 临界区代码
mutex.Unlock()

// sync.RWMutex（读写锁）
mutex.RLock()    // 读锁（多个 goroutine 可同时读）
mutex.RUnlock()
mutex.Lock()     // 写锁（互斥）
mutex.Unlock()

// sync.WaitGroup（等待一组 goroutine 完成）
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
}()
wg.Wait()
```

### 14.6 协程中的 panic 处理

每个 goroutine 需用 `defer + recover` 捕获 panic，否则会导致整个程序崩溃。

---

## 十五、反射（demo22）

```go
import "reflect"

// 获取类型
t := reflect.TypeOf(x)
t.Name()  // 类型名称（如 "Person"）
t.Kind()  // 底层类型（如 struct、int、slice）

// 获取值
v := reflect.ValueOf(x)

// 结构体字段
t.Field(0)           // 按索引获取字段
t.FieldByName("Age") // 按名称获取字段
t.NumField()         // 字段数量
v.FieldByName("Name") // 获取字段值
field.Tag.Get("json") // 获取 Tag

// 结构体方法
t.NumMethod()                     // 方法数量
v.MethodByName("Print").Call(nil) // 调用无参方法
v.MethodByName("SetInfo").Call([]reflect.Value{...}) // 调用有参方法
```

---

## 十六、文件操作（demo23）

```go
import "os"
```

| 操作 | 代码 |
|------|------|
| 打开文件(只读) | `os.Open("./main.go")` |
| 打开文件(可写) | `os.OpenFile("test.txt", os.O_CREATE\|os.O_RDWR, 0666)` |
| 读取文件 | `ioutil.ReadFile("./main.go")` |
| 写入文件 | `ioutil.WriteFile("test.txt", []byte(str), 0666)` |
| 创建目录 | `os.Mkdir("./abc", 0666)` |
| 创建多级目录 | `os.MkdirAll("dir1/dir2/dir3", 0666)` |
| 重命名 | `os.Rename("old.txt", "new.txt")` |
| 删除文件 | `os.Remove("t.txt")` |
| 删除目录 | `os.RemoveAll("aaa")` |

```go
// bufio 逐行读取
reader := bufio.NewReader(file)
line, err := reader.ReadString('\n')

// bufio 写入
writer := bufio.NewWriter(file)
writer.WriteString("你好golang\r\n")
writer.Flush()  // 必须 flush
```

---

## 十七、泛型（demo26，Go 1.18+）

### 17.1 泛型函数

```go
// 定义
func getData[T any](value T) T {
    return value
}

// 调用（类型可推断时可省略）
str1 := getData[string]("this is str")
num := getData(123)
```

### 17.2 类型约束

```go
// 自定义类型集接口
type Number interface {
    int | int64 | float64
}

func Add[T Number](a, b T) T {
    return a + b
}
```

### 17.3 泛型结构体

```go
type Container[T any] struct {
    value T
}

func (c *Container[T]) Set(value T) { c.value = value }
func (c Container[T]) Get() T       { return c.value }

intContainer := Container[int]{}
stringContainer := Container[string]{}
```

### 17.4 泛型接口

```go
type Usber[T any] interface {
    Start()
    Stop()
    GetDevice() T
}
```

---

## 十八、排序算法（demo11）

### 选择排序

每轮选出最小（或最大）值放在正确位置：

```go
for i := 0; i < len(numSlice); i++ {
    for j := i + 1; j < len(numSlice); j++ {
        if numSlice[i] > numSlice[j] {
            temp := numSlice[i]
            numSlice[i] = numSlice[j]
            numSlice[j] = temp
        }
    }
}
```

### 冒泡排序

相邻两两比较，每轮把最大（或最小）值"冒泡"到最后：

```go
for i := 0; i < len(numSlice); i++ {
    for j := 0; j < len(numSlice)-1-i; j++ {
        if numSlice[j] > numSlice[j+1] {
            temp := numSlice[j]
            numSlice[j] = numSlice[j+1]
            numSlice[j+1] = temp
        }
    }
}
```

### sort 包

```go
import "sort"

// 升序
sort.Ints(intList)
sort.Float64s(floatList)
sort.Strings(stringList)

// 降序
sort.Sort(sort.Reverse(sort.IntSlice(intList)))
```

---

## 十九、时间处理（demo14）

```go
import "time"

timeObj := time.Now()  // 当前时间

// 格式化(Go 诞生时间: 2006-01-02 15:04:05)
timeObj.Format("2006-01-02 15:04:05")  // 24小时制
timeObj.Format("2006-01-02 03:04:05")  // 12小时制

// 时间戳
timeObj.Unix()       // 秒时间戳
timeObj.UnixNano()   // 纳秒时间戳

// 时间戳转时间
time.Unix(int64(unixTime), 0)

// 字符串转时间
time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)

// 时间操作
timeObj.Add(time.Hour)  // 加1小时
time.Sleep(time.Second) // 休眠1秒

// 定时器
ticker := time.NewTicker(time.Second)
for t := range ticker.C { }
ticker.Stop()
```

---

## 二十、Redis 操作（demo24/demo25）

使用 `github.com/go-redis/redis/v8` 包：

```go
rdb := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
    DB:   0,
})

// 字符串
rdb.Set(ctx, "username", "zhangsan", 0)
rdb.Get(ctx, "username").Result()

// 列表
rdb.LPush(ctx, "hobby", "吃饭", "睡觉")
rdb.LRange(ctx, "hobby", 0, -1).Result()

// 集合
rdb.SAdd(ctx, "hobby", "吃饭", "睡觉")
rdb.SMembers(ctx, "hobby").Result()

// 哈希
rdb.HMSet(ctx, "userinfo", map[string]interface{}{"username": "张三"})
rdb.HGetAll(ctx, "userinfo").Result()

// 过期
rdb.Expire(ctx, "hobby", time.Second*20)

// 发布订阅
rdb.Publish(ctx, "server0", "message")
pubsub := rdb.Subscribe(ctx, "server0")
```

---

## 二十一、deciaml 精确计算（demo05）

用于解决 float 精度丢失问题：

```go
import "github.com/shopspring/decimal"

d1 := decimal.NewFromFloat(3.1).Add(decimal.NewFromFloat(4.2))
m3 := decimal.NewFromFloat(8.2).Sub(decimal.NewFromFloat(3.8))
// 还有 Mul(乘)、Div(除)
```

---

## 二十二、常见易错点总结

1. **变量必须使用**：定义后未使用会编译报错
2. **短变量 `:=`**：只能用于函数内，不能用于全局变量
3. **`{` 不能换行**：if/for/switch 的 `{` 必须和声明在同一行
4. **`++/--`**：只能独立使用，只有后置，且是语句不是表达式
5. **数组长度是类型的一部分**：`[3]int` 和 `[4]int` 不能混用
6. **slice/map/指针默认为 nil**：必须 make/new 初始化后才能使用
7. **Map 是无序的**：遍历顺序不确定
8. **defer 参数在注册时确定**：不是执行时
9. **结构体小写字段是私有的**：包外不能访问
10. **接口是隐式实现**：不需要显式声明 implements
11. **goroutine 需要 recover**：防止 panic 导致整个程序崩溃
12. **管道会阻塞**：满了不能写，空了不能读
