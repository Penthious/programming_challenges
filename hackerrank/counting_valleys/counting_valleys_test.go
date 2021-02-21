// https://www.hackerrank.com/challenges/counting-valleys/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
package hackerrank

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingValleys(t *testing.T) {
	assert.Equal(t, int32(1), countingValleys("UDDDUDUU"))
}

func countingValleys(steps string) int32 {
	var count int32 = 0
	var valleys int32 = 0
	sa := strings.Split(steps, "")
	inValley := false

	for _, step := range sa {
		if strings.ToLower(string(step)) == "u" {
			count++
		} else {
			count--
		}

		if count >= 0 && inValley {
			valleys++
			inValley = false
		} else if count < 0 {
			inValley = true
		}
	}

	return int32(valleys)
}
