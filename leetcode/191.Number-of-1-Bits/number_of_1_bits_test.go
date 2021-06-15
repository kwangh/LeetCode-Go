package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingWeight(t *testing.T) {
	a := assert.New(t)
	a.Equal(3, hammingWeight(11), "The input binary string 00000000000000000000000000001011 has a total of three '1' bits.")
	a.Equal(1, hammingWeight(64), "The input binary string 00000000000000000000000010000000 has a total of one '1' bit.")
	a.Equal(31, hammingWeight(4294967293), "The input binary string 11111111111111111111111111111101 has a total of thirty one '1' bits.")
}
