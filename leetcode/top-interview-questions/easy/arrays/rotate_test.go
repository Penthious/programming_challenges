package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {

	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3))
}
func rotate(nums []int, k int) []int { // T(1) S(n)
	nums = append(nums[len(nums)-k:], nums[:len(nums)-k]...)

	return nums
}
