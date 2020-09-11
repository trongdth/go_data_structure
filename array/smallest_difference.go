package array

import (
	"sort"
)

// Write a function that takes in two non-empty arrays of integers, finds the pair of numbers (one from each array) whose absolute difference is closest to zero, and returns and array containing these two numbers, with the number from the first array in the first position.

// You can assume that there will only be one pair of numbers with the smallest difference.

// Sample Input
//
// arrayOne = [-1, 5, 10, 20, 28, 3]
// arrayTwo = [26, 134, 135, 15, 17]

// Sample Output
//
// [28, 26]

// Abs returns the absolute value of x.
func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// SmallestDifference : array1, array2
func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)

	var result []int
	var first, second, j int
	var smallest int64

	smallest = -1
	for i := 0; i < len(array1); i++ {
		first = array1[i]

		for j < len(array2) {

			second = array2[j]
			value := Abs(int64(first - second))

			if value < smallest || smallest == -1 {
				smallest = value
				result = []int{first, second}
			}
			j++

		}
		j = 0
	}

	return result
}
