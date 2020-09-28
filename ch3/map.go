package main

import (
	"fmt",
	"sort"
)

func main() {
	// 使用make方法创建一个map的结构
	// 建立一个key类型为string，值类型为int的map
	// ages := make(map[string]int)

	// 使用字面量的方式创建map
	// ages := map[string]int{
	// 	"alice": 31,
	// 	"charlie": 34
	// }

	// 上面的创建方式和下面等价
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34


	// map的取值
	fmt.Println(ages["alice"]) // => 31

	// 从map中移除key value
	// delete(ages, "alice")

	// 在map中通过key取值的时候，当key不存在时则会返回值类型的zero type
	fmt.Println(ages["zhangjie"]) // => 0

	// map中的元素不是一个变量，所以我们不能直接去取地址
	// _ = &ages["alice"] // => compile error: cannot take address of map element
	// 不能取地址的一个原因是： 随着map里的数据增加，可能会rehashing已有的元素到一个新的地址，所以取地址的话会造成潜在的非法

	// 对map进行枚举
	// for name, age := range ages {
	// 	fmt.Printf("%s\t%d\n", name, age)
	// }

	// 使用string的sort方法对map进行排序
	// 拿到ages的keys，并且将其置入一个slice中，然后通过sort.Strings方法对slice进行排序
	// 拿到拍完序的slice，通过该顺序去map中取值，就可以达到顺序枚举map
	var names []string
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s\t\d\n", name, ages[name])
	}
}