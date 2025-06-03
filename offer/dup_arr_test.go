/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

// arr.len = n  0=<arr[i]<n 找重复的值
func DupArr(arr []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] == i { // 试图调整为 arr[i]=i 的样式
			continue
		}
		if arr[arr[i]] == arr[i] {
			res = append(res, arr[i])
		} else {
			arr[arr[i]], arr[i] = arr[i], arr[arr[i]]
		}
	}
	return res
}

func TestDupArr(t *testing.T) {
	fmt.Println(DupArr([]int{1, 2, 1}))
}
