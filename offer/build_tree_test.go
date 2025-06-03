/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(pre []int, mid []int) *TreeNode {
	if len(pre) == 0 || len(mid) == 0 {
		return nil
	}
	idx := findIdx(mid, pre[0])
	return &TreeNode{
		Val:   pre[0],
		Left:  BuildTree(pre[1:idx+1], mid[:idx]),
		Right: BuildTree(pre[idx+1:], mid[idx+1:]),
	}
}

func findIdx(arr []int, val int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return i
		}
	}
	return -1
}

func TestBuildTree(t *testing.T) {
	tree := BuildTree([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6})
	fmt.Println(tree)
}
