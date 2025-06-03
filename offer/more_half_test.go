/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

func MoreHalf(arr []int) int {
	res := 0
	cnt := 0
	for i := 0; i < len(arr); i++ {
		if cnt == 0 { // 为 0 的话随便跟一个
			res = arr[i]
			cnt++
			continue
		}
		if res == arr[i] { // 相同就增强，不同就削弱，超过半数的一定会增强到最后
			cnt++
		} else {
			cnt--
		}
	}
	return res
}

func TestMoreHalf(t *testing.T) {
	fmt.Println(MoreHalf([]int{1, 2, 3, 4, 5, 1, 1, 1, 1}))
}
