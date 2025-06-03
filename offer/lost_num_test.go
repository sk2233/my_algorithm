/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

// arr.len = n   0=<arr[i]<=n   只少一个数，找出来  只多一个数也可以这样玩
func LostNum(arr []int) int {
	res := (1 + len(arr)) * len(arr) / 2
	for i := 0; i < len(arr); i++ {
		res -= arr[i]
	}
	return res
}

func TestLostNum(t *testing.T) {
	fmt.Println(LostNum([]int{0, 1, 3}))
}
