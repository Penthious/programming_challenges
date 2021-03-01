package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	items := []int{44, 99, 6, 2, 1, 5, 63, 87, 283, 4, 0}

	assert.Equal(t, []int{0, 1, 2, 4, 5, 6, 44, 63, 87, 99, 283}, selectionSort(items))

}

func selectionSort(items []int) []int {
	type smallest struct {
		index int
		value int
	}
	counter := 0

	for counter < len(items) {
		s := &smallest{index: counter, value: items[counter]}
		for i := counter; i < len(items); i++ {
			if s.value > items[i] {
				s.value = items[i]
				s.index = i
			}

			if i+1 > len(items)-1 {
				items[counter], items[s.index] = s.value, items[counter]
			}
		}

		counter++
	}
	return items
}
