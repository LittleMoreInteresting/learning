package main

import (
	"fmt"
	"net/http"
)

func GetApi() {
	api := "http://localhost:8081/"

	res, err := http.Get(api)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		fmt.Printf("get api success\n")
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		GetApi()
	}
}