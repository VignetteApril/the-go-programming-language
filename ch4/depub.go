package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	map1 := map[string]int{
		"xx": 1,
	}

	map2 := map[string]int{
		"yy": 1,
	}

	map3 := map[string]int{
		"xx": 1,
	}

	fmt.Printf("%t\n", equal(map1, map2))

	fmt.Printf("%t\n", equal(map1, map3))

	seen := make(map[string]bool)
	// 这部分是如何获取从console输入的文字
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "depub: %v\n", err)
	}
	os.Exit(1)
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}