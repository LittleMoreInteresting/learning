package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	_ = r.ParseForm()
	form := r.Form          // 获取Map
	id := r.FormValue("id") // 获取单个参数
	fmt.Println(form)
	fmt.Println(string(body))

	//执行主要业务逻辑
	str := muxCurl()
	_, _ = fmt.Fprintln(w, "hello world", id, str)
}

func muxCurl() string {

	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	group, _ := errgroup.WithContext(timeout)
	var data1, data2 string

	group.Go(func() error {
		time.Sleep(10 * time.Second) // 设置超时模拟某个任务出错
		//time.Sleep(1 * time.Second)
		return nil
	})
	group.Go(func() error {
		fmt.Println("执行任务1")
		data1 = "任务1"
		return nil
	})
	group.Go(func() error {
		fmt.Println("执行任务2")
		data2 = "任务2"
		return nil
	})
	errCh := make(chan error, 1)
	errCh <- group.Wait()
	select {
	case <-timeout.Done():
		fmt.Println("超时 Back")
		return "err"
	case err := <-errCh:
		if err == nil {
			return fmt.Sprintf("%s;%s", data1, data2)
		} else {
			fmt.Printf("get error:%v", err)
			return ""
		}
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	_ = http.ListenAndServe("127.0.0.1:1234", nil)
}
