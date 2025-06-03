/*
@author: sk
@date: 2025/6/1
*/
package zhu_ji

import (
	"fmt"
	"testing"
)

// 0=< arr[i] <=8000 且都不重复
func BitSort(arr []int) []int {
	bits := make([]uint8, 1000)
	for i := 0; i < len(arr); i++ {
		bits[arr[i]/8] |= 1 << (arr[i] % 8)
	}
	res := make([]int, 0)
	for i := 0; i < 8000; i++ {
		if bits[i/8]&(1<<(i%8)) != 0 {
			res = append(res, i)
		}
	}
	return res
}

func TestBitSort(t *testing.T) {
	res := BitSort([]int{2, 4, 6, 3, 1})
	fmt.Println(res)
}
