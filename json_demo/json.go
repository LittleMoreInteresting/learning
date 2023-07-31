package main

import (
	"encoding/json"
	"fmt"
)

type NodeRole struct {
	Id   int64  `json:"id"`   // 角色ID
	Name string `json:"name"` // 角色名
}
type Node struct {
	Type     int       `json:"type"`
	Node     *NodeRole `json:"node"`
	Child    []*Node   `json:"child"`
	PartNode []*Node   `json:"partNode"`
}

func main() {
	str1 := "[{\"type\":1,\"node\":{\"name\":\"1\"},\"child\":[]},{\"type\":2,\"partNode\":[{\"type\":1,\"node\":{\"name\":\"21\"},\"child\":[{\"type\":1,\"node\":{\"name\":\"21\"},\"child\":[]}]},{\"type\":1,\"node\":{\"name\":\"22\"},\"child\":[{\"type\":2,\"partNode\":[{\"type\":1,\"node\":{\"name\":\"22\"},\"child\":[]},{\"type\":1,\"node\":{\"name\":\"22\"},\"child\":[]}]}]}]}]"
	var data []*Node
	err := json.Unmarshal([]byte(str1), &data)
	fmt.Printf("%v:%v \n", data, err)
	MoveTree(data)
}

func MoveTree(process []*Node) {
	if len(process) == 0 {
		return
	}
	for _, node := range process {
		if node.Type == 1 {
			fmt.Println(node.Node.Name)
			MoveTree(node.Child)
		} else if node.Type == 2 {
			MoveTree(node.PartNode)
		}
	}
}
