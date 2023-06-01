package main

import (
	"errors"
	"fmt"
	"sync"
)

func demo1() {
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
func demo2() {
	err_chen := make(chan error, 10)
	var err_num, success_num int
	var wg, wg1 sync.WaitGroup
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		for err := range err_chen {
			if err != nil {
				err_num++
			} else {
				success_num++
			}
		}
	}()

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			if i%2 == 0 {
				err_chen <- errors.New("error")
			} else {
				err_chen <- nil
			}
		}(i)
	}
	wg.Wait()
	close(err_chen)
	wg1.Wait()
	fmt.Println(err_num, success_num)
}
func main() {

	demo2()
}
