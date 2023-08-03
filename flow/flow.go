package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	json_data := loadJson()
	move, err := NewNodeMoveWithString(string(json_data))
	fmt.Println(err)
	var is_end bool
	var nodes = []*Node{}
	nodes, is_end = move.GetNextNodes(0)
	fmt.Println(is_end)
	for len(nodes) > 0 {
		node := nodes[0]
		fmt.Println(node.Node.Id)
		nodes = nodes[1:]
		//
		move.SetNodePass(node, &NodeRole{
			ApproveStatus: APPROVE_ACTION_PASS,
			ApproveOption: "tongg",
			ApproveUser:   "shr",
			ApproveTime:   time.Now().Unix(),
		})
		newNodes, isEnd := move.GetNextNodes(node.Id)
		nodes = append(nodes, newNodes...)
		fmt.Println(isEnd)
	}
	fmt.Println(move.ToString())
}

func loadJson() []byte {
	file, err := ioutil.ReadFile("D:\\code\\learning\\flow\\data.json")
	if err != nil {
		panic(err)
	}
	return file
}
