package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Children []Node `json:"children"`
}

func main() {
	str1 := "{\"id\":1,\"name\":\"node1\",\"children\":[{\"id\":2,\"name\":\"node2\",\"children\":[{\"id\":3,\"name\":\"node3\"}]}]}"
	var data Node
	err := json.Unmarshal([]byte(str1), &data)
	fmt.Printf("%v:%v \n", data, err)
	/*json_res, _ := json.Marshal(data)
	fmt.Printf("Json:%v\n", string(json_res))*/
}
