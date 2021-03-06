package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	// assert.Equal(t, []int{1, 2, 4}, plusOne([]int{1, 2, 3}))
	// assert.Equal(t, []int{4, 3, 2, 2}, plusOne([]int{4, 3, 2, 1}))
	// assert.Equal(t, []int{1, 0}, plusOne([]int{9}))
	assert.Equal(t,
		[]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 1, 0, 0},
		plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 9, 9}),
	)
	// assert.Equal(t, []int{1, 0, 0, 0, 0, 0}, plusOne([]int{9, 9, 9, 9, 9}))
	assert.Equal(t, []int{9, 0, 0, 0, 0}, plusOne([]int{8, 9, 9, 9, 9}))
}

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 == 10 {
			if i == 0 {
				digits[0] = 1
				digits = append(digits, 0)
			} else {
				digits[i] = 0
			}
		} else {
			digits[i]++
			break
		}

	}

	return digits
}
