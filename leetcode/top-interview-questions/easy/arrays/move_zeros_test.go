package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeros(t *testing.T) {

	assert.Equal(t, []int{1, 3, 12, 0, 0}, moveZeros([]int{0, 1, 0, 3, 12}))
	assert.Equal(t, []int{1, 0, 0}, moveZeros([]int{0, 0, 1}))
	assert.Equal(t, []int{4, 2, 4, 3, 5, 1, 0, 0, 0, 0}, moveZeros([]int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}))
}

func moveZeros(nums []int) []int {
	counter := 0
	index := 0
	for counter < len(nums) {
		if nums[index] == 0 {
			nums = append(nums[:index], nums[index+1:]...)
			nums = append(nums, 0)
		} else {
			index++
		}

		counter++
	}
	return nums
}
