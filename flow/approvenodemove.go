package main

import (
	"encoding/json"
)

var (
	APPROVE_ACTION_PASS   = 1
	APPROVE_ACTION_REJECT = 2
)

type NodeRole struct {
	Id            int64  `json:"id"`   // 角色ID
	Name          string `json:"name"` // 角色名
	ApproveStatus int    `json:"approve_status"`
	ApproveUser   string `json:"approve_user"`
	ApproveTime   int64  `json:"approve_time"`
	ApproveOption string `json:"approve_option"`
}
type Node struct {
	Id            int       `json:"id"`
	Type          int       `json:"type"`
	Node          *NodeRole `json:"node"`
	Child         []*Node   `json:"child"`
	PartNode      []*Node   `json:"partNode"`
	Pre           *Node     `json:"-"`
	Next          *Node     `json:"-"`
	ApproveStatus int       `json:"approve_status"`
}

type NodeMove struct {
	flowNodes []*Node
}

func NewNodeMove(flowNodes []*Node) *NodeMove {
	buildTree(flowNodes, nil)
	return &NodeMove{
		flowNodes: flowNodes,
	}
}
func NewNodeMoveWithString(nodeData string) (*NodeMove, error) {
	var flowNodes []*Node
	err := json.Unmarshal([]byte(nodeData), &flowNodes)
	buildTree(flowNodes, nil)
	return &NodeMove{
		flowNodes: flowNodes,
	}, err
}

// 获取下一批审批节点 , 是否审批完成
func (move *NodeMove) GetNextNodes(role_id int64) ([]*Node, bool) {
	if role_id == 0 {
		node := move.flowNodes[0]
		if node.Type == 1 {
			return []*Node{node}, false
		} else if node.Type == 2 {
			//return node.PartNode, false
			nextNewNode := []*Node{}
			list := node.PartNode
			for len(list) > 0 {
				nextNode := list[0]
				list = list[1:]
				if nextNode.Type == 1 {
					nextNewNode = append(nextNewNode, nextNode)
				} else if nextNode.Type == 2 {
					list = append(list, nextNode.PartNode...)
				}
			}
			return nextNewNode, false
		}
	}
	node := move.SearchNodeByRoleId(role_id)

	// 有串联下级
	if len(node.Child) > 0 {
		nextNewNode := []*Node{}
		list := node.Child
		for len(list) > 0 {
			nextNode := list[0]
			list = list[1:]
			if nextNode.Type == 1 {
				nextNewNode = append(nextNewNode, nextNode)
			} else if nextNode.Type == 2 {
				list = append(list, nextNode.PartNode...)
			}
		}
		return nextNewNode, false
	}
	// 并联子集
	for node.Pre != nil {
		if node.Pre.ApproveStatus == 1 {
			node = node.Pre
		} else {
			return []*Node{}, false
		}
	}
	if node.Next != nil {
		node = node.Next
		if node.Type == 1 {
			return []*Node{node}, false
		} else if node.Type == 2 {
			//return node.PartNode, false
			nextNewNode := []*Node{}
			list := node.PartNode
			for len(list) > 0 {
				nextNode := list[0]
				list = list[1:]
				if nextNode.Type == 1 {
					nextNewNode = append(nextNewNode, nextNode)
				} else if nextNode.Type == 2 {
					list = append(list, nextNode.PartNode...)
				}
			}
			return nextNewNode, false
		}
	}
	return []*Node{}, true
}

func (move *NodeMove) SearchNodeByRoleId(role_id int64) *Node {
	var found *Node
	list := move.flowNodes
	for len(list) > 0 {
		node := list[0]
		list = list[1:]
		if node.Type == 1 {
			if node.Node.Id == role_id {
				found = node
				return found
			}
			if len(node.Child) > 0 {
				list = append(list, node.Child...)
			}
		} else {
			list = append(list, node.PartNode...)
		}
	}
	return found
}

// 节点审批通过
func (move *NodeMove) SetNodePass(node *Node, role *NodeRole) {
	if node.Type == 1 {
		node.ApproveStatus = APPROVE_ACTION_PASS
		node.Node.ApproveStatus = role.ApproveStatus
		node.Node.ApproveUser = role.ApproveUser
		node.Node.ApproveTime = role.ApproveTime
		node.Node.ApproveOption = role.ApproveOption

		if len(node.Child) == 0 && node.Pre != nil {
			move.SetNodePass(node.Pre, role)
		}
		return
	}
	if node.Type == 2 {
		list := node.PartNode
		allPass := true
		for len(list) != 0 {
			n := list[0]
			list = list[1:]
			if n.ApproveStatus == 0 {
				allPass = false
				break
			}
			if len(n.Child) != 0 {
				list = append(list, n.Child...)
			}
			if len(n.PartNode) != 0 {
				list = append(list, n.PartNode...)
			}
		}
		if allPass {
			node.ApproveStatus = 1
			if len(node.Child) == 0 && node.Pre != nil {
				move.SetNodePass(node.Pre, role)
			}
		}
	}
}

// 节点审批驳回
func (move *NodeMove) SetNodeReject(node *Node, role *NodeRole) {
	node.ApproveStatus = APPROVE_ACTION_REJECT
	node.Node.ApproveStatus = role.ApproveStatus
	node.Node.ApproveUser = role.ApproveUser
	node.Node.ApproveTime = role.ApproveTime
	node.Node.ApproveOption = role.ApproveOption
}

func buildTree(process []*Node, pre *Node) {
	if len(process) == 0 {
		return
	}
	var current *Node
	for i, _ := range process {
		node := process[i]
		if node.Type == 1 {
			buildTree(node.Child, pre)
		} else if node.Type == 2 {
			buildTree(node.PartNode, node)
		}
		if pre != nil {
			node.Pre = pre
		} else {
			if current != nil {
				current.Next = node
			}
		}
		current = node
	}
}

func (move *NodeMove) ToString() string {
	marshal, _ := json.Marshal(move.flowNodes)
	return string(marshal)
}
