/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

type Node struct {
	Val  int
	Next *Node
}

// 给出要删除的节点，改节点非头节点或尾节点，删除它
func NodeDelete(node *Node) {
	// 拷贝下一个节点到当前节点并删除下一个节点
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func TestNodeDelete(t *testing.T) {
	root := &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}}
	NodeDelete(root.Next)
	fmt.Println(root)
}
