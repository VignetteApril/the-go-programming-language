package main

import (
	"fmt"
)

// Array  的长度是固定的，也是因为固定长度数组基本上不会go编程时直接使用
// Slices 长度可以自由伸缩，所以更常用一些在go中
// 但是想要了解Slices之前，要先了解Array
func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a) - 1])

	// 遍历数组，输出index 和 值
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// 遍历数组，只输出值
	for _, v := range a{
		fmt.Printf("%d\n", v)
	}

	// 使用array的字面量形式创建一个数组
	// var q [3]int = [3]int{ 1, 2, 3 }
	var r [3]int = [3]int{ 1, 2 }
	fmt.Println(r[2]) // => 0 当取值超过数组的长度会自动取0值
}

// 补充下关于print的相关知识
// Printf 全程是 print formatter 可以在在字符串内使用各种占位符，然后再在后面传入相关的值
// Print 无法format任何东西，单纯的输出字符串
// Println 全称 print line 和print方法相同，但是会在末尾加上换行符
