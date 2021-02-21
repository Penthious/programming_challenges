/*
Given a m * n matrix of ones (representing soldiers) and zeros (representing civilians),
return the indexes of the (k) weakest rows in the matrix ordered from the weakest to the strongest.

A row i is weaker than row j, if the number of soldiers in row i is less than the number of
soldiers in row j, or they have the same number of soldiers but i is less than j. Soldiers are
always stand in the frontier of a row, that is, always ones may appear first and then zeros.

Constraints:

m == mat.length
n == mat[i].length
2 <= n, m <= 100
1 <= k <= m
matrix[i][j] is either 0 or 1.
*/
package leetcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeakestRows(t *testing.T) {
	matrix1 := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	matrix2 := [][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}
	matrix3 := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{1, 1, 1},
		{1, 1, 1},
		{0, 0, 0},
		{1, 1, 1},
		{1, 0, 0},
	}

	expectedOutput := []int{2, 0, 3, 1, 4}
	assert.Equal(t, expectedOutput, kWeakestRows(matrix1, 5))

	expectedOutput = []int{0, 2}
	assert.Equal(t, expectedOutput, kWeakestRows(matrix2, 2))

	expectedOutput = []int{4, 6, 0, 1, 2, 3}
	assert.Equal(t, expectedOutput, kWeakestRows(matrix3, 6))
}

func kWeakestRows(matrix [][]int, k int) []int {
	var output []int
	set := make(map[int]int)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 {
				set[i] = 0
			}
			if matrix[i][j] == 1 {
				set[i] = set[i] + 1
			}
		}
	}

	type data struct {
		Key   int
		Value int
	}
	var d []data

	for i := 0; i < len(set); i++ {
		d = append(d, data{i, set[i]})
	}

	sort.Slice(d, func(i, j int) bool {
		if d[i].Key < d[j].Key && d[i].Value == d[j].Value {
			return true
		}
		return d[i].Value < d[j].Value
	})

	for i := 0; i < k; i++ {
		output = append(output, d[i].Key)
	}

	return output
}
