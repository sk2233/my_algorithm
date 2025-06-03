/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

// 0123456789  10111213   0~n
// 100 101
// 0000000000  1 1 1 1 1 1 1   2 2 2 2 2 2
func OneCount(n int) int {
	n++
	// 个位  每 10 个 1个
	// 十位  每 100个 10个
	// 百位  每 1000个 100个
	res := 0
	base1 := 10
	base2 := 1
	for {
		res += n / base1 * base2
		if n%base1 > base2 {
			res += min((n%base1)-base2, base2)
		}
		if n/base1 == 0 {
			break
		}
		base1 *= 10
		base2 *= 10
	}
	return res
}

func TestOneCount(t *testing.T) {
	fmt.Println(OneCount(10))
}
