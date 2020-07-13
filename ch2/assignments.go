// 赋值
package main

import (
	"fmt"
)

func main() {
	var x int
	x = 1

	var p_b bool
	p = &p_b
	*p = true

}

// 间接赋值
// 1、方法的参数会被间接赋值，根据传进来的参数
// 2、方法的return语句，会把return的values赋值到对应返回值列表上
// 3、对对数组中的值的赋值也是间接的
// 4、和数组类似，maps channels也同样是间接赋值

medals := []strings{ "gold", "silver", "bronze" }
medals[0] = "gold"
medals[1] = "silver"
medals[2] = "bronze"
