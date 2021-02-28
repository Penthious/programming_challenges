package coding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciRecursive(t *testing.T) {
	assert.Equal(t, 21, fibonacciRecursive(8))
	assert.Equal(t, 144, fibonacciRecursive(12))
}

func fibonacciRecursive(index int) int {

	if index < 2 {
		return index
	}

	return fibonacciRecursive(index-1) + fibonacciRecursive(index-2)
}
