package main

import (
	"flag"
	"fmt"
	"strings"
)

// flag.Bool 创建了一个新的flag变量类型为bool
// 这个方法接受三个参数：
// 1、flag的名称，本例子中为 n
// 2、这个变量的默认值，本例子为 false
// 3、如果用户在执行命令行时加了-n则为true 不加则为false
var n = flag.Bool("n", false, "omit trailing newline")
// 同理flag.String 接受了一个名称，一个默认值，一个输出的信息，来创建一个string类型的flag变量
// 如果用户在执行命令行时输入了-s则会将-s后的字符作为新的分割符，如果不输入则使用默认的“ ”作为分隔符
var sep = flag.String("s", " ", "separator")

func main() {
	// 当程序运行时，需要先调用flag.Parse，然后再使用flags更新参数的值
	// 非flag参数可以从flag.Args()中取出，它的数据结构是字符串slice
	// 如果flag.Parse()出错，则会打印出错误信息，并调用os.Exit(2) 终止程序
	flag.Parse()
	fmt.Println(*n)
	fmt.Println(*sep)
	// 用flag包内方法生成的变量都是指针类型的（指向参数变量），所以必须用*操作符取出原有的值
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
