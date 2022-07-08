package main

import (
	"encoding/json"
	"fmt"
)

type Member struct {
	Id int `json:"id"`
	Name string `json:"name"`
}
func main() {
	member := &Member{
		1,"Tommy",
	}
	marshal, err := json.Marshal(member)
	if err!=nil {
		panic(err)
	}
	fmt.Printf("JSON:%s",marshal)
	mem2 := Member{}
	err = json.Unmarshal(marshal, &mem2)
	if err!=nil {
		panic(err)
	}
	fmt.Printf("MEM:%v",mem2)
}
