/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

type CNode struct {
	Val   int
	Next  *CNode
	Other *CNode // 旁路节点
}

func CNodeCopy(node *CNode) *CNode {
	// 先拷贝一份放在老的后面
	curr := node
	for curr != nil {
		curr.Next = &CNode{
			Val:  curr.Val,
			Next: curr.Next,
		}
		curr = curr.Next.Next
	}
	// 建立连接
	curr = node
	for curr != nil {
		if curr.Other != nil {
			curr.Next.Other = curr.Other.Next
		}
		curr = curr.Next.Next
	}
	// 拆分链表
	curr = node
	newCurr := curr.Next
	res := newCurr
	for curr.Next.Next != nil {
		curr.Next = curr.Next.Next
		curr = curr.Next
		newCurr.Next = curr.Next
		newCurr = newCurr.Next
	}
	return res
}

func TestCNodeCopy(t *testing.T) {
	c1 := &CNode{Val: 1}
	c2 := &CNode{Val: 2}
	c3 := &CNode{Val: 3}
	c4 := &CNode{Val: 4}
	c5 := &CNode{Val: 5}
	c1.Next = c2
	c1.Other = c3
	c2.Next = c3
	c2.Other = c5
	c3.Next = c4
	c4.Next = c5
	c4.Other = c2
	node := CNodeCopy(c1)
	fmt.Println(node)
}
