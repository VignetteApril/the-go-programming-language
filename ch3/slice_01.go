package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("cap: ", cap(s), "len: ", len(s))

	reverse(s[:2])
	fmt.Println(s)
	reverse(s[2:])
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)
	// 4.2 Un like arrays, slices are not comparable, so we cannot use == to test whether two slices cont ain
	// the same elements.
	// 数组可以通过 == 进行比较
	// 切片不能通过 == 进行比较
	// 特例1：[]byte 这种类型的切片可以通过标准库提供的bytes.Equal 方法比较
	// 其他的类型的slice则需要通过手动实现比对方法
	// Why?

	// The only legal slice comp arison is against niL
	// 唯一能使用 == 比较slice的是nil


	// 4.2.1 append function
	var runes []rune

	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}

	fmt.Printf("%q\n", runes)
}

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}



