/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

// arr 中有一个数只出现了一次，其他都出现 3 次，找出来
func OnlyOne(arr []int) int {
	cnts := make([]int, 32) // 收集所有数字，所有位上的 01 和
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		idx := 0
		for temp > 0 {
			if temp&1 == 1 {
				cnts[idx]++
			}
			temp >>= 1
			idx++
		}
	}
	res := 0
	base := 1
	for i := 0; i < len(cnts); i++ {
		if cnts[i]%3 == 1 { // 出现 3 次的肯定没有余数，只有那个出现一次的会有余数  出现 5 次啥的也是类似
			res += base
		}
		base <<= 1
	}
	return res
}

func TestOnlyOne(t *testing.T) {
	fmt.Println(OnlyOne([]int{1, 1, 1, 2, 3, 3, 3}))
}
