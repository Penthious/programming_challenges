package coding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecursionFactorial(t *testing.T) {
	assert.Equal(t, 120, recursionFactorial(5))
	assert.Equal(t, 3628800, recursionFactorial(10))
}

func recursionFactorial(number int) int {
	if number == 1 {
		return 1
	}
	return number * recursionFactorial(number-1)
}
