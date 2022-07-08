package main

import "fmt"

func main() {
	/*ch := make(chan int,1)

	go func() {
		time.Sleep(3*time.Second)
		ch <- 1
	}()

	after := time.After(2 * time.Second)

	select {
	case <-ch:
		fmt.Println("get data from ch")
	case <-after:
		fmt.Println("time out")
	}*/
	conns := []int{1,2,3,4,5,6,7,8,9,10}
	ch := make(chan int,1)
	for _, conn := range conns {
		go func(conn int) {
			select {
			case ch <- conn:
			default:
				fmt.Println("default")
			}
		}(conn)
	}
	c := <-ch
	fmt.Println(c)
}
