/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

func OnlyOne2(arr []int) (int, int) {
	// 得到那两个不同值的异或值，其他相同的都异或为 0 了
	temp := 0
	for i := 0; i < len(arr); i++ {
		temp ^= arr[i]
	}
	// 找那两个值异或为 1 的位置 (这个位置它们俩肯定是一个 1 一个 0)
	base := 1
	for temp&base == 0 {
		base <<= 1
	}
	res1, res2 := 0, 0
	for i := 0; i < len(arr); i++ {
		if arr[i]&temp == 0 { // 使用这个特殊位分为两个数组每个各一个单一值  只有一个单一值直接异或即可
			res1 ^= arr[i]
		} else {
			res2 ^= arr[i]
		}
	}
	return res1, res2
}

func TestOnlyOne2(t *testing.T) {
	fmt.Println(OnlyOne2([]int{2, 4, 3, 6, 3, 2, 5, 5}))
}
