package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPrimes(t *testing.T) {
	a := assert.New(t)
	a.Equal(4, countPrimes(10), "There are 4 prime numbers less than 10, they are 2, 3, 5, 7.")
	a.Equal(0, countPrimes(0))
	a.Equal(0, countPrimes(1))
	a.Equal(0, countPrimes(2))
}
