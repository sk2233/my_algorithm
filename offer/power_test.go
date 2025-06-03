/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

func Power(val float64, n int) float64 {
	res := 1.0
	base := val
	for n > 0 {
		if n&1 == 1 {
			res *= base
		}
		n /= 2
		base *= base
	}
	return res
}

func TestPower(t *testing.T) {
	fmt.Println(Power(2, 5))
}
