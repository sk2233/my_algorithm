/*
@author: sk
@date: 2025/6/1
*/
package zhu_ji

import (
	"fmt"
	"testing"
)

// 0<= arr[i] <=1000 其值互不相同，其少了至少一个 0～1000 的值找出来一个
func TwoPoints(arr []int) int {
	l, r := 0, 1000
	for l < r {
		m := (l + r) / 2
		if count(arr, l, m) < m-l+1 {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func count(arr []int, l int, r int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] >= l && arr[i] <= r {
			res++
		}
	}
	return res
}

func TestTwoPoints(t *testing.T) {
	fmt.Println(TwoPoints([]int{1, 23, 4, 0}))
}
