// There is a large pile of socks that must be paired by color. Given an array of integers representing the color of each sock, determine how many pairs of socks with matching colors there are.
package hackerrank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfPairs(t *testing.T) {
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}

	assert.Equal(t, int32(3), sockMerchant(ar), "Sock pairs dont match")
}

func sockMerchant(ar []int32) int32 {
	pairs := make(map[int32]int)
	var output int32

	for _, value := range ar {
		pairs[value]++
		if pairs[value]%2 == 0 {
			output++
			pairs[value] = 0

		}
	}
	return output
}
