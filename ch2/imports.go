package main

import (
	"fmt"
	"os"
	"strconv"
	"gopl.io/ch2/tempconv" // 导入自己创建的包的时候需要把包放置在$GOPATH的src里面，
												 // 并且自己要创建一个同名的文件夹，然后把同名的go文件放置进去
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
