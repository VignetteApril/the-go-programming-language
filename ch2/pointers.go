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

	// 如果使用指针作为参数传进方法里，这样就可以间接的修改指针所指向的变量的值

	v := 1
	incr(&v) // => 2
	fmt.Println(incr(&v))
}

func incr(p *int) int {
	*p++
	return *p
}


// var p = f()

// 这里返回的是变量v的地址指针，每当程序调用该方法完毕后返回变量v的地址指针，但是变量v并没有
// 被销毁，依然被指针p饮用着
// 疑问点：如果没有指针饮用的变量是否被垃圾回收了呢
// func f() *int {
// 	v := 1
// 	return &v
// }

// 每次调用返回的地址都是会变的，说明每次调用方法 f 都会创建一个新的变量 v，该变量会拥有
// 新的地址
// fmt.Println(f() == f())
// => always false
// because



