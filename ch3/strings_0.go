package main

import "fmt"

func main() {
	s := "hello world"
	fmt.Println(len(s)) // => 11 输出11个字节大小

	s = "hello world文"
	fmt.Println(len(s)) // => 14 这里输出的是14，因为在uft-8中汉字占3或4个字节

	c := s[11]
	fmt.Println(c) // 通过索引取字符串在该索引字符的ASCII码
	fmt.Println(string([]rune(s)[0])) // => utf-8 h 通过这样的方式可以取出字符串的utf-8格式的
}
