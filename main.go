package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Unix() / 100 * 100
	fmt.Println(now)
	fmt.Println(time.Unix(now, 0).Format("2006-01-02 15:04:05"))
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
