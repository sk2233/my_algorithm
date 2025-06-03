/*
@author: sk
@date: 2025/6/3
*/
package offer

import (
	"fmt"
	"testing"
)

func BuyStock(arr []int) int {
	minVal := arr[0]
	res := 0
	for i := 1; i < len(arr); i++ {
		if arr[i]-minVal > res {
			res = arr[i] - minVal
		}
		if arr[i] < minVal { // 只要保留前序的最小值即可
			minVal = arr[i]
		}
	}
	return res
}

func TestBuyStock(t *testing.T) {
	fmt.Println(BuyStock([]int{9, 11, 8, 5, 7, 12, 16, 14}))
}
