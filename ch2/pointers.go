// Pointers
package main

import (
	"fmt"
)

func main() {
	// 指针数据类型
	// 声明一个叫 x 类型为 int 初始值为 1 的变量
	x := 1

	// 通过 & 操作符获取变量的指针
	p := &x
	// 通过 * 操作符获取指针指向的值
	fmt.Println(*p)

	// 通过对 *p 进行赋值则会改变变量 x 的值
	*p = 2
	fmt.Println(x)

	// Println formats using the default formats for its operands and writes to
	// standard output. Spaces are always added between operands and a newline is
	// appended. It returns the number of bytes written and any write error encountered.

	var j, k int
	fmt.Println(&j == &j, &j == &k, &j == nil)
	// => true false false

	// The zero value for a pointer of any type is nil
	// 疑问点：如何声明一个zero value 的 pointer
}
