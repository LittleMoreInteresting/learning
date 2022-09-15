package main

import (
	"context"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)

var ctx = context.Background()
var esUrl = "http://192.168.72.128:9201/"
var EsClient *elastic.Client

// 初始化es连接
func InitEs() {
	// 连接es客户端
	client, err := elastic.NewClient(
		elastic.SetURL(esUrl),
	)
	if err != nil {
		log.Fatal("es 连接失败:", err)
	}
	info, code, err := client.Ping(esUrl).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Elasticsearch returned with code>: %d and version %s\n", code, info.Version.Number)

	EsClient = client
	fmt.Println("es连接成功")
}

func main() {
	InitEs()
}

/***
elastic.SetSniff(false)
no active connection found: no Elasticsearch node available
*/
