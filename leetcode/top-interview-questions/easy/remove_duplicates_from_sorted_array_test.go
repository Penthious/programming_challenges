package easy

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/727/
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	assert.Equal(t, []int{0, 1, 2, 3, 4}, removeDuplicatesFromSortedArray(nums))
}

func removeDuplicatesFromSortedArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	working := true
	index := 0

	for working {
		working = false
		if index+1 >= len(nums) {
			return nums
		}

		if nums[index] == nums[index+1] {
			nums = append(nums[:index], append([]int{}, nums[index+1:]...)...)
			working = true
		} else {
			index++
			working = true
		}

	}
	return nums
}
