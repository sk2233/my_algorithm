/*
@author: sk
@date: 2025/6/1
*/
package zhu_ji

import (
	"fmt"
	"testing"
)

func MaxSub(arr []int) int {
	res := 0
	temp := 0
	for i := 0; i < len(arr); i++ {
		if temp+arr[i] > 0 { // 子串只要还大于 0 就能使用
			temp += arr[i]
		} else { // <=0 就直接抛弃掉
			temp = 0
		}
		if temp > res {
			res = temp
		}
	}
	return res
}

func TestMaxSub(t *testing.T) {
	fmt.Println(MaxSub([]int{1, 2, -3, 4, 2}))
}
