/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

// 012345678910
// 10  90  900  9000
func NumOfSeq(num int) int {
	base := 1
	cnt := 10
	for {
		if num > base*cnt {
			num -= base * cnt
		} else {
			break
		}
		if base == 1 {
			cnt = 9
		} else {
			cnt *= 10
		}
		base++
	}
	temp := 0
	for i := 1; i < base; i++ {
		temp = temp*10 + 9
	}
	temp += (num+base-1)/base - 1
	idx := base - (num % base)
	for i := 1; i < idx; i++ {
		temp /= 10
	}
	return temp % 10
}

func TestNumOfSeq(t *testing.T) {
	fmt.Println(NumOfSeq(12))
}
