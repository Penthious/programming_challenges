package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	items := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}

	assert.Equal(t, []int{0, 1, 2, 4, 5, 6, 44, 63, 87, 99, 283}, bubbleSort(items))

}

func bubbleSort(items []int) []int {
	working := true
	for working {
		working = false
		for i := 1; i < len(items); i++ {
			if items[i-1] > items[i] {
				items[i], items[i-1] = items[i-1], items[i]
				working = true
			}
		}
	}

	return items
}
