package main

import "fmt"

func moveZeroes(num [] int) [] int {
	l,r,n := 0,0,len(num)
	for r < n {
		if num[r] != 0 {
			num[r],num[l] = num[l],num[r]
			l++
		}
		r++
	}
	return num;
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Printf("%v", nums)
}
