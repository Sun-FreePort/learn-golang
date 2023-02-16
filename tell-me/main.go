package main

import (
	"fmt"
	"os"
)

func Hello() string {
	return "Hi，欢迎你来到这里，我见过你吗？似乎有些眼熟。\n我是这个世界的管理者，你想知道什么？你可以输入这些数字获得信息：\n - 1. 背包\n - 2. 身体\n - 3. 新闻\n - 0. 退出"
}

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}