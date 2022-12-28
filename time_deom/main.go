package main

import (
	"fmt"
	"time"
)

const (
	DateTime = "2006-01-02 15:04:05"
	DateOnly = "2006-01-02"
	TimeOnly = "15:04:05"
)

func main() {
	now := time.Now()
	fmt.Println("当前时间戳：", now.Unix())
	fmt.Println("当前时间：", now.UTC())
	fmt.Println("格式化1：", now.Format(DateTime))
	fmt.Println("格式化2：", now.Format(DateOnly))
	fmt.Println("格式化3：", now.Format(TimeOnly))

	//指定时区
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	fmt.Println("SH : ", time.Now().In(cstSh).Format(DateTime))

}
