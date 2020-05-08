package main

// Write a function that takes in a non-empty array of distinct integers and an
// integer representing a target sum. If any two numbers in the input array sum
// up to the target sum, the function should return them in an array, in any
// order. If no two numbers sum up to the target sum, the function should return
// an empty array.

// sample input
// 	array = [3, 5, -4, 8, 11, 1, -1, 6]
// 	target sum = 10

// sample output
// 	[-1, 10]

// TwoNumberSum : array, target
func TwoNumberSum(array []int, target int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if i != j {
				if array[i]+array[j] == target {
					return []int{array[i], array[j]}
				}
			}
		}
	}

	return nil
}
