/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

func UglyNum(num int) int {
	arr := make([]int, num)
	arr[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < num; i++ {
		minNum := arr[i2] * 2
		if arr[i3]*3 < minNum {
			minNum = arr[i3] * 3
		}
		if arr[i5]*5 < minNum {
			minNum = arr[i5] * 5
		}
		arr[i] = minNum
		if arr[i2]*2 == minNum {
			i2++
		}
		if arr[i3]*3 == minNum {
			i3++
		}
		if arr[i5]*5 == minNum {
			i5++
		}
	}
	fmt.Println(arr)
	return arr[num-1]
}

func TestUglyNum(t *testing.T) {
	fmt.Println(UglyNum(10))
}
