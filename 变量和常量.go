package main

import "fmt"

// 演示不同的变量声明方式
var name string = "小明" // 标准方式
var age = 25           // 类型推断
var (
	height    float64 = 175.5 // 身高
	weight    int     = 70    // 体重
	isStudent bool    = false // 是否为学生
)

// 包级常量声明
const Pi = 3.14159
const (
	StatusOK    = 200
	StatusError = 500
)

// 使用iota的枚举示例
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	// 函数内的变量声明
	salary := 8000 // 短变量声明

	// 演示变量的零值
	var defaultInt int
	var defaultString string
	var defaultBool bool

	fmt.Println("变量演示:")
	fmt.Printf("姓名: %s, 年龄: %d\n", name, age)
	fmt.Printf("身高: %.1f, 体重: %d, 是否学生: %v\n", height, weight, isStudent)
	fmt.Printf("月薪: %d\n", salary)

	fmt.Println("\n零值演示:")
	fmt.Printf("整数默认值: %d\n", defaultInt)
	fmt.Printf("字符串默认值: %q\n", defaultString)
	fmt.Printf("布尔值默认值: %v\n", defaultBool)

	fmt.Println("\n常量演示:")
	fmt.Printf("圆周率: %f\n", Pi)
	fmt.Printf("状态码: %d, %d\n", StatusOK, StatusError)

	fmt.Println("\n星期枚举:")
	fmt.Printf("Sunday: %d, Monday: %d\n", Sunday, Monday)

	// 演示作用域
	{
		innerVar := "我是代码块内的变量"
		fmt.Println(innerVar)
	}
	// 这里已经不能访问 innerVar 了
}
