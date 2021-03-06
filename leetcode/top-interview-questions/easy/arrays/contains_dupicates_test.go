package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicates(t *testing.T) {
	assert.Equal(t, true, containsDuplicates([]int{1, 2, 3, 1}))
	assert.Equal(t, false, containsDuplicates([]int{1, 2, 3, 4}))
	assert.Equal(t, true, containsDuplicates([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

func containsDuplicates(list []int) bool {
	hashTable := make(map[int]bool)

	for _, v := range list {
		if hashTable[v] != false {
			return true
		}
		hashTable[v] = true
	}

	return false
}
