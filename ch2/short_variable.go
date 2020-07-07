package main

import (
	"fmt"
)

// var 声明和 := 声明的使用场景不区别
// var 1、声明的变量需要一个准确的类型
//     2、声明变量后面还要对他进行赋值
// 其他的变量声明都可以用 := 的方式进行声明


// 疑问点：tuple assignment
// i, j = j, i 这样的操作是如何实现的

// 注意点1：一个short variable declaration 不需要在左侧声明所有的变量，如果有一些变量已经
// 在同一个作用域已经声明过了的话则short variable declaration 表象的就像赋值
var l int

func main(){
	i := 100
	j := 200

	fmt.Printf("i = %d, j = %d\n", i, j)
	// => i = 100, j = 200

	i, j = j, i
	fmt.Printf("i = %d, j = %d\n", i, j)
	// => i = 200, j = 100


	// short variable declaration 中左侧有已声明过的变量时
	i, k := 300, 100
	fmt.Printf("i = %d, k = %d\n", i, k)
	// => i = 300, k = 100

	// short variable declaration 中不能全都是已声明过的变量，必须要有一个是新声明的
	// i, j := 100, 200
	// => ./short_variable.go:37:7: no new variables on left side of :=

	// short variable declaration 表现为赋值时只针对当前作用域的，作用域外的变量则不会起作用
	// 如果l被 := 认为是已经声明的变量时，则下面的语句会报出no new variables on left side of :=
	// 但是其实是不会报错的，因为short variable declaration只认当前作用域的变量
	i, l := 200, 200
	fmt.Printf("i = %d, l = %d\n", i, l)
	// => i = 200, l = 200
}



