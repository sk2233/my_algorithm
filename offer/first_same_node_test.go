/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

func FirstSameNode(node1, node2 *Node) *Node {
	cnt1 := 0
	curr := node1
	for curr != nil {
		curr = curr.Next
		cnt1++
	}
	cnt2 := 0
	curr = node2
	for curr != nil {
		curr = curr.Next
		cnt2++
	}
	if cnt1 < cnt2 {
		cnt1, cnt2 = cnt2, cnt1
		node1, node2 = node2, node1
	}
	for i := 0; i < cnt1-cnt2; i++ {
		node1 = node1.Next
	}
	for node1 != node2 {
		node1 = node1.Next
		node2 = node2.Next
	}
	return node1
}

func TestFirstSameNode(t *testing.T) {
	same := &Node{Val: 22, Next: &Node{Val: 33}}
	node1 := &Node{Val: 1, Next: &Node{Val: 2, Next: same}}
	node2 := &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: same}}}
	res := FirstSameNode(node1, node2)
	fmt.Println(res)
}
