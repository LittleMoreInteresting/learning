package main

import (
	"fmt"
	"strings"
)

func main() {
	split := strings.Split("abc(kdkkd)", "(")
	fmt.Println(split[0])
}

func strLen(str []byte) int {
	n := len(str)
	for i := n - 1; i >= 0; i-- {
		if i != n-1 && string(str[i]) == " " {
			return n - 1 - i
		}
	}
	return n
}
