package main

import (
	"fmt"
	"crypto/sha256"
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

	// 在申请数组长度的位置使用...，则可以根据数组的字面量长度确定数组的长度
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // => "[3]int"

	var name = "Kim"
	var age = 22

	fmt.Print(name, " is ", age, " years old.\n")  // Kim is 22 years old.
	fmt.Printf("%v is %v years old.\n", name, age) // Kim is 22 years old.
	fmt.Println(name, "is", age, "years old.")     // Kim is 22 years old.

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBR
		RMB
	)

	// 这种形式的方式是将index显示的显现出来
	// 比如USD就是1 EUR就是2等等
	symbol := [...]string{USD: "$", EUR: "xx", GBR: "xxx", RMB: "XXXX"}
	fmt.Println(RMB, symbol[RMB]) // => "3 XXXX"

	// 如果数组里的元素是可以比较的，则该数组就是可以比较的
	// 所以我们是可以直接比较数组的
	e := [2]int{1, 2}
	f := [...]int{1, 2}
	g := [2]int{1, 3}
	fmt.Println(e == f, f == g, e == g)
	// fmt.Println(a == d)
	// => comiple error: cannot compare: 知识点：因为数组的长度也
	// 是type的一部分，所以长度不同的数组属于不同type的数据类型，所以不能比较


	// crypto sha256
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	var array_pointer [32]byte = [32]byte{1}
	normal_zero(array_pointer)
	fmt.Println(array_pointer)
}

// 补充下关于print的相关知识
// Printf 全程是 print formatter 可以在在字符串内使用各种占位符，然后再在后面传入相关的值
// Print 无法format任何东西，单纯的输出字符串
// Println 全称 print line 和print方法相同，但是会在末尾加上换行符
// var name = "Kim"
// var age = 22

// fmt.Print(name, " is ", age, " years old.\n")  // Kim is 22 years old.
// fmt.Printf("%v is %v years old.\n", name, age) // Kim is 22 years old.
// fmt.Println(name, "is", age, "years old.")     // Kim is 22 years old.


// 定义一个方法，接受一个指针参数（类型为[32]byte）
// 该参数就是该指针指向的那个变量，在方法内可以对那个变量进行修改，只有这种方法可以在方法内修改变量的值
// 如果不是指针参数，则参数列表中的值只是传入参数值的引用，并不能直接改变变量的值
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

// 如上述所示，该方法不能改变原变量的值
func normal_zero(ptr [32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

// // Exercise 4.1
// // Write a function that counts the number of bits that are different in two SHA256 hashes
// func compare_sha256(first_hash []byte, second []byte) {
// 	return len(first_hash)
// }
