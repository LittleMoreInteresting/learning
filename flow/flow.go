package main

import (
	"fmt"
	"os"
)

func main() {
	json_data := loadJson()
	move, err := NewNodeMoveWithString(string(json_data))
	fmt.Println(err)

	var is_end bool
	var nodes = []*Node{}
	nodes, is_end = move.GetNextNodes(8)
	fmt.Println(nodes[0])
	/*for len(nodes) > 0 {
		node := nodes[0]
		fmt.Println(node.Id, node.Node.Id, node.Node.Name)
		nodes = nodes[1:]

		move.SetNodePass(node, &NodeRole{
			ApproveStatus: APPROVE_ACTION_PASS,
			ApproveOption: "tongg",
			ApproveUser:   "shr",
			ApproveTime:   time.Now().Unix(),
		})
		var newNodes []*Node
		newNodes, is_end = move.GetNextNodes(node.Id)

		nodes = append(nodes, newNodes...)

	}*/
	//fmt.Println(move.ToString())
	fmt.Println(is_end)
}

func loadJson() []byte {
	file, err := os.ReadFile("D:\\code\\learning\\flow\\data.json")
	if err != nil {
		panic(err)
	}
	return file
}
