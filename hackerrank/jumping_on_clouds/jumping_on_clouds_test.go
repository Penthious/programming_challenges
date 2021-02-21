package hackerrank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpingOnClouds(t *testing.T) {
	clouds := []int32{0, 0, 1, 0, 0, 1, 0, 0, 1, 0}
	// 0 0 1
	// 0 1 0
	// 0 0 1
	// 0 1 0
	// 0 0 1
	// 0 1 0
	clouds1 := []int32{0, 0, 0, 1, 0, 0}

	assert.Equal(t, int32(6), jumpingOnClouds(clouds))
	assert.Equal(t, int32(3), jumpingOnClouds(clouds1))
}

func jumpingOnClouds(c []int32) int32 {
	var count int32 = 0

	for i := 0; i < len(c); i++ { // o(n)
		// contains := make(map[int]int)
		// if len(c) > i+2 && c[i] ==0 {
		// 	contains[i] =
		// }
		if len(c) > i+2 && contains(c[i:i+3], 1) && c[i] == 0 { // o(n^2) redo to figure out how to do O(n)
			count++
		} else if i != len(c)-1 && !contains(c[i:], 1) {
			count++
			i++
		}

	}

	return count
}

func contains(c []int32, value int32) bool {
	for _, v := range c {
		if v == value {
			return true
		}
	}
	return false
}
