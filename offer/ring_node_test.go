/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

func RingNode(node *Node) *Node {
	fast := node.Next.Next
	slow := node.Next
	for fast != slow {
		fast = fast.Next.Next
		slow = slow.Next
	}
	res := node
	for slow != res {
		res = res.Next
		slow = slow.Next
	}
	return res
}

func TestRingNode(t *testing.T) {
	node := &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: nil}}}}}
	node.Next.Next.Next.Next.Next = node.Next
	res := RingNode(node)
	fmt.Println(res)
}
