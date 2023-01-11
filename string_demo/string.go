package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	str := "你好"
	fmt.Printf("%v\n", len(str))
	fmt.Printf("%v\n", utf8.RuneCountInString(str))
	fmt.Println(strings.ToTitle("hello world"))
	fmt.Println(strings.ToUpper("hello world"))
	fmt.Println(Ucfirst("hello world"))
}
func Ucfirst(str string) (uc string) {
	for i, c := range str {
		uc = string(unicode.ToUpper(c)) + str[i+1:]
		return
	}
	return
}
