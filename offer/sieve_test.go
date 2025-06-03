/*
@author: sk
@date: 2025/6/3
*/
package offer

import (
	"fmt"
	"testing"
)

func Sieve(n, s int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, s) // i j  i+1 个筛子 搞出 j+1 的种类数
	}
	for i := 0; i < s; i++ {
		if i+1 <= 6 {
			dp[0][i] = 1
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j < s; j++ { // i j  i+1 个筛子 搞出 j+1 的种类数
			for k := 1; k <= 6 && j-k >= 0; k++ { // 当前从 1 到 6
				dp[i][j] += dp[i-1][j-k]
			}
		}
	}
	return dp[n-1][s-1]
}

//  1 2 3 4 5 6

func TestSieve(t *testing.T) {
	fmt.Println(Sieve(2, 3))
}
