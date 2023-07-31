package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4}
	for len(list) != 0 {
		i := list[0]
		list = list[1:]
		fmt.Println(i)
		if i == 3 {
			list = append(list, 5, 6, 7)
		}
	}
}
