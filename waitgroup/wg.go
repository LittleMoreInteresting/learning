package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		var s []int
		for i := 0; i < 10; i++ {
			s = append(s, i)
			fmt.Println(len(s), cap(s), s)
		}
	}()
	wg.Wait()
	fmt.Println("ok")
}
