/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

// arr.len = n    0<arr[i]<n  必定有重复值，找一个
func CountDup(arr []int) int {
	l, r := 1, len(arr)-1
	for l < r {
		m := (l + r) / 2
		if count(arr, l, m) > m-l+1 {
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

func TestCountDup(t *testing.T) {
	fmt.Println(CountDup([]int{1, 2, 3, 1}))
}
