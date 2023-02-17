package main

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/hashicorp/consul/api"
	"learning/consul/conf"

	"github.com/go-kratos/kratos/contrib/config/consul/v2"
)

func main() {
	consulClient, err := api.NewClient(&api.Config{
		Address: "172.16.1.100:8500",
	})
	if err != nil {
		panic(err)
	}
	cs, err := consul.New(consulClient, consul.WithPath("dev/bss/pay-by-stage/service"))
	// consul中需要标注文件后缀，kratos读取配置需要适配文件后缀
	// The file suffix needs to be marked, and kratos needs to adapt the file suffix to read the configuration.
	if err != nil {
		panic(err)
	}
	c := config.New(config.WithSource(cs))
	if err = c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err = c.Scan(&bc); err != nil {
		panic(err)
	}
	fmt.Println(bc)
}
