/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"testing"
)

func FirstSingleChar(str string) byte {
	temp := make(map[byte]int)
	arr := []byte(str)
	for i, b := range arr {
		if _, ok := temp[b]; ok {
			temp[b] = -1 // 再次出现摧毁下标
		} else {
			temp[b] = i // 第一次出现记录下标
		}
	}
	res := byte(0)
	minIdx := len(str)
	for k, v := range temp {
		if v >= 0 && v < minIdx { // 找没有摧毁且下标最靠前的
			minIdx = v
			res = k
		}
	}
	return res
}

func TestFirstSingleChar(t *testing.T) {
	fmt.Println(string([]byte{FirstSingleChar("google")}))
}
