package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	// assert.Equal(t, 1, singleNumber([]int{2, 2, 1}))
	assert.Equal(t, 4, singleNumber([]int{4, 1, 2, 1, 2}))
	// assert.Equal(t, 1, singleNumber([]int{1}))
}

func singleNumber(nums []int) int { // T(n) S(1)
	output := 0

	for _, v := range nums {
		output ^= v
	}
	return output
}
