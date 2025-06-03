/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

// 寻找逆对数 抽象为归并排序
func ReverseCount(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	mid := len(arr) / 2
	res := ReverseCount(arr[:mid]) // 先收集子集的逆对数
	res += ReverseCount(arr[mid:])
	temp := make([]int, len(arr))
	li, ri := 0, mid
	i := 0
	for li < mid && ri < len(arr) {
		if arr[li] > arr[ri] {
			temp[i] = arr[li]
			i++
			li++
			res += len(arr) - ri // 左边每有一个领先 就大于全部右边剩余的，这些都是逆对数
		} else {
			temp[i] = arr[ri]
			i++
			ri++
		}
	}
	for li < mid {
		temp[i] = arr[li]
		i++
		li++
	}
	for ri < len(arr) {
		temp[i] = arr[ri]
		i++
		ri++
	}
	copy(arr, temp)
	return res
}

func TestReverseCount(t *testing.T) {
	fmt.Println(ReverseCount([]int{7, 5, 6, 4}))
}
