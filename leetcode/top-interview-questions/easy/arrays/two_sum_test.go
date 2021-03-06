package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 2}, twoSum([]int{3, 2, 4}, 6))
	assert.Equal(t, []int{0, 2}, twoSum([]int{3, 2, 4}, 7))
}

func twoSum(nums []int, target int) []int {
	ht := make(map[int]int)

	for i, v := range nums {
		if ht[v] != 0 {
			return []int{ht[v] - 1, i}
		}

		ht[target-v] = i + 1

	}

	return []int{}
}
