package coding

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseAStringRecursive(t *testing.T) {
	reversed := "yllib"
	fmt.Println(reverseAStringRecursive(reversed))
	assert.Equal(t, "billy", reverseAStringRecursive(reversed))

}
func reverseAStringRecursive(input string) string {
	data := []rune(input)
	if len(data) == 1 {
		return string(data[0])
	}
	output := data[len(data)-1]
	return fmt.Sprintf("%s%v", string(output), reverseAStringRecursive(string(data[:len(data)-1])))
}
