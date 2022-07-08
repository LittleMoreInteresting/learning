package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Database string `mapstructuor:"dabases"`
	Host string `mapstructuor:"host"`
	Port int64 `mapstructuor:"port"`
	Username string `mapstructuor:"username"`
	Password string `mapstructuor:"password"`
}

func main() {
	v := viper.New()
	v.SetConfigFile("viper/config.yml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	name := v.Get("name")
	dbConfig := MysqlConfig{}
	err = v.UnmarshalKey("Mysql",&dbConfig)

	if err != nil {
		panic(err)
	}
	fmt.Printf("config.name=%v",name)
	fmt.Printf("config.mysql=%v",dbConfig)
}
