/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

func TopK(arr []int, k int) []int {
	if len(arr) < 2 {
		return arr[:k]
	}
	mid := arr[0]
	l, r := 0, len(arr)-1
	for l < r {
		for arr[r] >= mid && l < r {
			r--
		}
		if l == r {
			break
		}
		arr[l] = arr[r]
		for arr[l] <= mid && l < r {
			l++
		}
		if l == r {
			break
		}
		arr[r] = arr[l]
	}
	arr[l] = mid
	if l > k {
		TopK(arr[:l], k)
	}
	if l+1 < k {
		TopK(arr[l+1:], k-l-1)
	}
	return arr[:k]
}

func TestTopK(t *testing.T) {
	fmt.Println(TopK([]int{3, 5, 6, 1, 2}, 2))
}
