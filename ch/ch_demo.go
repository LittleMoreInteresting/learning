package main

import "fmt"

func main() {

	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	/*for num := range ch {
		fmt.Println(num)
	}*/
	for {
		num, ok := <-ch
		if !ok {
			fmt.Println("channel closed")
			return
		}
		fmt.Println(num)
	}
}
