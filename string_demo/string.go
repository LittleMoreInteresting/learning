package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "你好";
	fmt.Printf("%v\n",len(str))
	fmt.Printf("%v\n",utf8.RuneCountInString(str))
}
