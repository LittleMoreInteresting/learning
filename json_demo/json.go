package main

import (
	"encoding/json"
	"fmt"
)

type DoubleSlices struct {
	Str     []string `json:"str"`
	Numbers []int    `json:"numbers"`
}

func main() {
	ds := &DoubleSlices{
		[]string{"a", "b", "c", "d"},
		[]int{1, 2, 3, 4},
	}
	marshal, err := json.Marshal(ds)
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON:%s", marshal)
	mem2 := DoubleSlices{}
	err = json.Unmarshal(marshal, &mem2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("MEM:%v", mem2)
}
