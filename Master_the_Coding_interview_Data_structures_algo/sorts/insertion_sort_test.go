package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	items := []int{44, 99, 6, 2, 1, 5, 63, 87, 283, 4, 0}

	assert.Equal(t, []int{0, 1, 2, 4, 5, 6, 44, 63, 87, 99, 283}, insertionSort(items))

}

func insertionSort(items []int) []int {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
	return items
}
