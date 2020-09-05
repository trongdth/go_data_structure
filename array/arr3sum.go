package array

import (
	"sort"
)

// Write a function that takes in a non-empty array of distinct integers and an
// integer representing a target sum. The function should find all triplets in
// the array that sum up to the target sum and return a two-dimensional array of
// all these triplets. The numbers in each triplet should be ordered in ascending
// order, and the triplets themselves should be ordered in ascending order with
// respect to the numbers they hold.

// sample input
// 	array = [12, 3, 1, 2, -6, 5, -8, 6]
// 	target sum = 0

// sample output
// 	[[-8, 2, 6], [-8, 3, 5], [-6, 1, 5]]

// ThreeNumberSum : array, target
func ThreeNumberSum(array []int, target int) [][]int {

	sort.Ints(array)
	result := make([][]int, 0)
	var value, left, right int
	for i := 0; i < len(array)-2; i++ {
		value = array[i]
		left = i + 1
		right = len(array) - 1
		for left < right {
			sum := value + array[left] + array[right]

			if sum == target {
				result = append(result, []int{value, array[left], array[right]})
				left++
			} else {
				if sum < target {

					if array[left+1] > array[left] {
						left++
					}

				} else if sum > target {

					if array[right-1] < array[right] {
						right--
					}

				}
			}

		}
	}

	return result
}
