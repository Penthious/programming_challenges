package coding

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseAstring(t *testing.T) {

	word := "esreveR"

	assert.Equal(t, "Reverse", reverseWord(word), "The word was not reversed")
}

func reverseWord(word string) string {
	w := strings.Split(word, "")
	for i := len(w)/2 - 1; i >= 0; i-- {
		opp := len(w) - 1 - i
		w[i], w[opp] = w[opp], w[i]
	}

	return strings.Join(w, "")
}
