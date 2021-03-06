package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersectionOfTwoArrays(t *testing.T) {
	assert.Equal(t, []int{2, 2}, intersectionOfTwoArrays([]int{1, 2, 2, 1}, []int{2, 2}))
}

func intersectionOfTwoArrays(nums1, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		return intersectionOfTwoArrays(nums2, nums1)
	}

	ht := make(map[int]int)
	output := []int{}

	for _, v := range nums1 {
		ht[v]++
	}

	for _, v := range nums2 {
		if ht[v] != 0 {
			output = append(output, v)
			ht[v]--
		}
	}

	return output
}
