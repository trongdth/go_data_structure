package array

// Write a function that takes in an array of integers and returns the length of the longest peak in the
// array.

// A peak is defined as adjacent integers in the array that are strictly increasing until they reach a tip
// (the highest value in the peak), at which point they become strictly decreasing. At least three
// integers are required to form a peak.

// For example, the integers 1, 4, 10, 2 form a peak, but the integers 4, 0, 10 don't and
// neither do the integers 1, 2, 2, 0. Similary, the integers 1, 2, 3 don't form a peak because
// there aren't any strictly decreasing integers after the 3.

// Sample Input
// array = [1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3]

// Sample Output
// 6 [0, 10, 6, 5, -1, -3]

// LongestPeak : array
func LongestPeak(array []int) int {
	var longestLen, index int

	index = 1
	for index < len(array)-1 {
		if array[index-1] < array[index] && array[index] > array[index+1] {

			var left, right int
			left = index - 1
			for left >= 0 && array[left+1]-array[left] > 0 {
				left--
			}
			right = index + 1
			for right <= len(array)-1 && array[right-1]-array[right] > 0 {
				right++
			}

			tmp := (index - left) + (right - index) + 1 - 2
			if longestLen < tmp {
				longestLen = tmp
			}
		}
		index++
	}

	return longestLen

}
