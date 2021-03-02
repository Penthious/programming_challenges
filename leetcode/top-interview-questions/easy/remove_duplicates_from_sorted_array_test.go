package easy

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/727/
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	removeDuplicatesFromSortedArray(nums)
	assert.Equal(t, 4, len(nums))
	assert.Equal(t, []int{0, 1, 2, 3, 4}, nums)
}

func removeDuplicatesFromSortedArray(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	working := true
	index := 0

	for working {
		working = false
		if index+1 >= len(nums) {
			return len(nums)
		}

		if nums[index] == nums[index+1] {
			nums = append(nums[:index], append([]int{}, nums[index+1:]...)...)
			working = true
		} else {
			index++
			working = true
		}

	}
	return len(nums)
}
