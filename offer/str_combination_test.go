/*
@author: sk
@date: 2025/6/2
*/
package offer

import (
	"fmt"
	"sort"
	"testing"
)

func Combination(arr string) {
	bs := []byte(arr)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})
	for i := 0; i < len(bs); i++ {
		Build([]byte{bs[i]}, bs, i+1)
	}
}

func Build(temp []byte, bs []byte, i int) {
	if i >= len(bs) {
		fmt.Println(string(temp))
		return
	}
	newTmep := make([]byte, len(temp))
	copy(newTmep, temp)
	Build(temp, bs, i+1)
	newTmep = append(newTmep, bs[i])
	Build(newTmep, bs, i+1)
}

func TestCombination(t *testing.T) {
	Combination("abc")
}
