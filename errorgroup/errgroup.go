package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func worker(ctx context.Context, eg *errgroup.Group,ch chan int,id int) {
	eg.Go(func() error {
		<-ctx.Done()
		fmt.Println("Do Back")
		return nil
	})

	eg.Go(func() error {
		for i := range ch {
			if i == 8 {
				return errors.New("Bet ID")
			}
			fmt.Printf("[%v]:Do something %v……\n",id,i)
		}
		return nil
	})
}

func main() {
	//ctx,cancel := context.WithCancel(context.Background())
	group, errCtx := errgroup.WithContext(context.Background())
	ch := make(chan int,10)
	group.Go(func() error {
		for i := 0; i < 10; i++ {
			ch<- i
			time.Sleep(1*time.Second)
		}
		close(ch)
		return errors.New("End")
	})
	worker(errCtx,group,ch,1)
	worker(errCtx,group,ch,2)


	err := group.Wait()
	if err == nil {
		fmt.Println("都完成了")
	} else {
		fmt.Printf("get error:%v", err)
	}

}
