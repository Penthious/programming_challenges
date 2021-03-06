package arrays

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidSudoku(t *testing.T) {

	board1 := [][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
	board2 := [][]string{
		{"8", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
	board1Bytes := [][]byte{}
	board2Bytes := [][]byte{}
	for _, v := range board1 {
		board1Bytes = append(board1Bytes, convertStringToBytes(v))
	}
	for _, v := range board2 {
		board2Bytes = append(board2Bytes, convertStringToBytes(v))
	}

	assert.Equal(t, true, validSudoku(board1Bytes))
	// assert.Equal(t, false, validSudoku(board2Bytes))

}
func convertStringToBytes(b []string) []byte {
	buf := &bytes.Buffer{}

	gob.NewEncoder(buf).Encode(b)
	return buf.Bytes()
}

func validSudoku(board [][]byte) bool {
	hs := make(map[string]bool)
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			var b []string
			buf := &bytes.Buffer{}
			dec := gob.NewDecoder(buf)
			dec.Decode(&b)

			// b := board[r][:]
			fmt.Println(b)
			if hs[fmt.Sprintf("byte %v %d row", b[c], r)] {

				return false
			}
			if hs[fmt.Sprintf("byte %v %d column", b[c], c)] {

				return false
			}
			if hs[fmt.Sprintf("byte %v box pos: %d %d", b[c], r/3, c/3)] {
				return false
			}
			if b[c] != "." {
				hs[fmt.Sprintf("byte %v %d row", b[c], r)] = true
				hs[fmt.Sprintf("byte %v %d column", b[c], c)] = true
				hs[fmt.Sprintf("byte %v box pos: %d %d", b[c], r/3, c/3)] = true
			}
		}
	}

	fmt.Println(hs)
	return true
}
