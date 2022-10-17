package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	HashList []int               `json:"hash_list"`
	TimeList []int               `json:"time_list"`
	Mp       []map[string]string `json:"mp"`
}

func main() {
	str1 := "{\"hash_list\":[12,12,12,12,12]}"
	str2 := "{\"time_list\":[34,34,34,34,34]}"
	var data Data
	_ = json.Unmarshal([]byte(str1), &data)
	_ = json.Unmarshal([]byte(str2), &data)
	fmt.Printf("H:%v\n", data.HashList)
	fmt.Printf("T:%v\n", data.TimeList)
	json_res, _ := json.Marshal(data)
	fmt.Printf("Json:%v\n", string(json_res))
}
