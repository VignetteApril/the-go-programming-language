package main

import "fmt"

// 不要把生命周期和域的概念搞混
// 声明周期是 运行时的概念
// 域是      编译时的概念
func main() {
	x := "hello"

	for _, x := range x {
		// 当这里重新声明的时候，当前的块会屏蔽块外的变量x
		// 相当于重新声明了一个x，而且不会对外面的变量x有影响
		x := x + 'A' - 'a'
		fmt.Printf("%c\n", x)
	}

	// char + 'A' - 'a' 貌似是go里面变成大写的一种机制
	// NOTE: 搞清楚内部原理
	y := 'h' + 'A' - 'a'
	fmt.Printf("%c\n", y)

	fmt.Printf("%s\n", x)
}
