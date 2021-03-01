package coding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type cache struct {
	Calculations int
	Data         map[int]int
}

func TestFibonacciCache(t *testing.T) {
	cache := &cache{}
	cache.Data = make(map[int]int)
	fib := fibonacciCache(10, cache)
	assert.Equal(t, 55, fib)

	assert.Equal(t, 9, cache.Calculations)
}

func fibonacciCache(n int, c *cache) int {
	if c.Data[n] != 0 {
		return c.Data[n]
	}

	if n < 2 {
		return n
	}
	c.Calculations++
	c.Data[n] = fibonacciCache(n-1, c) + fibonacciCache(n-2, c)
	return c.Data[n]
}
