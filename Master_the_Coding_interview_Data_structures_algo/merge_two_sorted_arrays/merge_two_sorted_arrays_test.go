package coding

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoSortedArrays(t *testing.T) {
	sorted1 := []int{0, 2, 4, 5}
	sorted2 := []int{1, 3, 4}
	assert.Equal(t, []int{0, 1, 2, 3, 4, 4, 5}, mergeTwoSortedArrays(sorted1, sorted2))
}

func mergeTwoSortedArrays(arr1, arr2 []int) []int {
	combined := append(arr1, arr2...)
	sort.Ints(combined)

	return combined
}
