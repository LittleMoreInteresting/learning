package main

import "fmt"

func main() {

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()
	/*for num := range ch {
		fmt.Println(num)
	}*/
	for {
		num, ok := <-ch1
		num2, ok2 := <-ch2
		if !ok || !ok2 {
			fmt.Println("channel closed")
			return
		}
		fmt.Println(num, num2)
	}
}
