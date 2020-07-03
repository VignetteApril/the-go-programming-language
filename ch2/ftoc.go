package main

import "fmt"


// 变量篇：
// Declaration example: var name type = expression
// type 和 = expression 可以省略其中一个，但不能都省略
// 当type省略时：变量的类型则会根据初始值来决定
// 当= expression省略时：则会被赋予一个 zero type value
// 当type为numbers初始值为：0
// 当type为boolean初始值为：false
// 当type为strings初始值为：""
// nil for type 为interfaces、reference types（slice, pointer, map, channel, function）
// zero-value 机制保证变量永远有一个well-defined值，所以在go中永远没有未初始化值的变量


func main() {
	const freezingF, boilingF	= 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToc(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToc(boilingF))
}

func fToc(f float64) float64 {
	return (f - 32) * 5 /9
}
