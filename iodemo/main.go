package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func showMenu() {
	fmt.Println("----IO Demo----")
	fmt.Println("1 表示 标准输入")
	fmt.Println("2 表示 普通文件")
	fmt.Println("3 表示 从字符串")
	fmt.Println("4 表示 从网络")
	fmt.Println("b 返回上级菜单")
	fmt.Println("q 退出")
	fmt.Println("***********************************")
	fmt.Println("Please input code：")

}

func main() {
	for {
		showMenu()
		var code string
		_, err := fmt.Scanln(&code)
		if err != nil {
			log.Fatal(err)
			return
		}
		switch strings.ToLower(code) {
		case "1":
			fmt.Println("code:"+code)

		case "q":
			fmt.Println("exit")
			os.Exit(0)
		}
	}
}
