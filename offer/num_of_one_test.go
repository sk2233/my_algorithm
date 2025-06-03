/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

func NumOfOne(num int) int {
	res := 0
	for num != 0 {
		num &= num - 1
		res += 1
	}
	return res
}

func TestNumOfOne(t *testing.T) {
	fmt.Println(NumOfOne(2233))
}
