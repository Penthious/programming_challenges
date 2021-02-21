/*
Given 2 arrays, create a function that
lets a user know (true/false) whether
these two arrays contain any common items.
*/
package coding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoArraysContainCommonItems(t *testing.T) {

	list1 := []string{"a", "b", "c", "d"}
	list2 := []string{"x", "y", "z"}
	list3 := []string{"m", "n", "d", "s"}

	assert.Equal(t, true, commonArrayChecker(list1, list3))
	assert.Equal(t, false, commonArrayChecker(list1, list2))
}

func commonArrayChecker(base []string, common []string) bool {

	exists := make(map[string]bool)

	for _, item := range base { // O(n)
		exists[item] = true
	}

	for _, item := range common { // O(n)
		if exists[item] {
			return true
		}
	}

	return false
}
