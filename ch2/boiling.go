package main

import "fmt"

// go 中的声明主要有4种：var, const, type, func
// var 声明一个变量，可以被改变
// const 声明一个常量，不能被改变
// func 声明一个方法
// type ?

// 该常量声明在‘包’级别的作用域
const boilingF = 212.0

// func 的声明包含几个部分 名称、参数列表、返回值type列表（可选）
func main() {
	// 变量 f、c 是在 方法 main里面这个作用域中的
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
