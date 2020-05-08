package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	assert := assert.New(t)

	array := []int{12, 3, 1, 2, -6, 5, -8, 6}
	targetSum := 0

	result := ThreeNumberSum(array, targetSum)
	assert.Equal([][]int{{-8, 2, 6}, {-8, 3, 5}, {-6, 1, 5}}, result)
}
