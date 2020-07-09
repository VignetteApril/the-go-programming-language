package main

import (
	"fmt"
)

func main() {
	// 表达吃 new(T)创建了一个未命名的变量，并且类型为T，并初始化这个类型的zero type
	// 然后 return 这个未命名变量的地址
	p := new(int)
	fmt.Println(*p)

	*p = 2
	fmt.Println(*p)
}
