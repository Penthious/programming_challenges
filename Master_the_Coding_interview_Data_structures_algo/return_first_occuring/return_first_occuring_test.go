package coding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnFirstOccuring(t *testing.T) {
	data := []int{2, 1, 3, 4, 2, 3}

	assert.Equal(t, 2, returnFirstOccuring(data))
}

func returnFirstOccuring(arr []int) int {
	m := make(map[int]bool)

	for _, value := range arr {
		if m[value] {
			return value
		}
		m[value] = true
	}

	return -1 // assumes all values are positive, if we want to include negative then instead of returning one value return int & err
}
