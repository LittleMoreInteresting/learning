package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func main() {
	/*
	   DialTimeout time.Duration `json:"dial-timeout"`
	   Endpoints []string `json:"endpoints"`
	*/
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("连接成功")
	defer cli.Close()
	//---------------------------------

	//设置1秒超时，访问etcd有超时控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	cancel()

	//取值，设置超时为1秒
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	cli.Put(ctx, "/message", "Hello1")
	resp, err := cli.Get(ctx, "/message")
	cli.Delete(ctx, "hello")
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
