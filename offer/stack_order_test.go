/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

func StackOrder(push, pop []int) bool {
	stack := make([]int, 0)
	popIdx := 0
	for i := 0; i < len(push); i++ {
		stack = append(stack, push[i])
		for len(stack) > 0 && popIdx < len(pop) && stack[len(stack)-1] == pop[popIdx] {
			stack = stack[:len(stack)-1]
			popIdx++
		}
	}
	return len(stack) == 0
}

func TestStackOrder(t *testing.T) {
	fmt.Println(StackOrder([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}
