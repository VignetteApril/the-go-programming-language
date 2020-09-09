package main

import (
	"fmt"
)

func main() {
	// my_slice := make([]int, 3, 5)
	// fmt.Println(len(my_slice))
	// fmt.Println(cap(my_slice))
	// fmt.Println("my slice:", my_slice)
	// println(my_slice) // => [3/5]0xc00001e0c0 这其实是slice的本质 其中3为长度 5为最大伸缩值 0xADDR为指向底层数组的地址
	// 这里也可以这么理解，slice为一种包含有另外两个值的不纯粗指针

	// 初始化、创建、访问slice
	// 1. 使用make()
	// 创建一个length 和 capacity 都为5的slice
	// slice := make([]int, 5)
	// 创建一个length为3 capacity 为 5的slice
	// slice := make([]int, 3, 5)

	// make 和 new 方法在底层的区别
	// make()比new()函数多一些操作，new()函数只会进行内存分配并做默认的赋0初始化，而make()可以
	// 先为底层数组分配好内存，然后从这个底层数组中再额外生成一个slice并初始化。另外，make只能构
	// 建slice、map和channel这3种结构的数据对象，因为它们都指向底层数据结构，都需要先为底层数据
	// 结构分配好内存并初始化。

	// 通过直接赋值的方式创建slice
	// 创建长度和容量都为4，并初始化赋值
	// color_slice := []string{"red", "blue", "black", "green"}

	// 创建长度和容量都为100的slice，并将第100个元素赋值为3
	// slice := []int{99: 3}

	// 注意区分数组和切片
	// 数组的创建在[]中有值，一般的为数字，或者...等
	// 创建长度为3的数组
	// array := [3]int{10, 20, 30}
	// 创建长度和容量都为3的切片
	// slice := []int{10, 20, 30}

	// 由于slice的底层是数组，可以通过索引的方式修改slice中元素的值
	// my_slice := []int{11, 22, 33, 44, 55}
	// fmt.Println(my_slice[2])
	// my_slice[2] = 333
	// fmt.Println(my_slice[2])


	// 对slice进行切片
	// SLICE[A:B:C] A表示从第几个元素开始切，B表示切到第几个元素 切片的长度 = B - A ，
	// C控制切片的容量，容量 = C - A， 如果被省略则默认是底层数组的最尾部
	my_slice := []int{11, 22 ,33 ,44 ,55}
	// 从第二个元素开始切，长度为2，容量为4
	new_slice := my_slice[1:3]

	// 由于多个切片共享一个底层数组，当有一个切片修改了底层数组的值时，其他的切片则会一起变化
	// 但是这样会造成混乱，因为在代码开发过程中我们无法保证哪个切片和哪个切片共享一个底层数组
	// 这时候我们需要copy函数进行处理
	// func copy(dst, src []Type) int
}
