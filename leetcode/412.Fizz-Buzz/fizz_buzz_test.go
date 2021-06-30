package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	a := assert.New(t)
	a.Equal([]string{"1", "2", "Fizz"}, fizzBuzz(3))
	a.Equal([]string{"1", "2", "Fizz", "4", "Buzz"}, fizzBuzz(5))
	a.Equal([]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}, fizzBuzz(15))
}
