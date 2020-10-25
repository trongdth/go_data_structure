package array

// Write a function that takes in an n x m two dimensional array (that can be square-shaped when
// n == m) and returns a one-dimensional array of all the array's elements in spiral order.

// Spiral order starts at the top left corner of the two-dimentsional array, goes to the right, and
// proceeds in a spiral pattern all the way until every element has been visited.

// Sample Input
//
// array = [
// 	[1,  2,  3,  4],
//  [12, 13, 14, 5],
//  [11, 16, 15, 6],
//  [10, 9, 8, 7],
// ]

// Sample Ouput
//
// [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]

// SpiralTraverse : array
func SpiralTraverse(array [][]int) []int {

	var arr []int
	var traversedArr [][]bool
	var left, top, right, bottom int
	bottom = len(array) - 1
	if bottom > 0 {
		right = len(array[0]) - 1
	}

	traversedArr = make([][]bool, len(array))
	for i := range array {
		traversedArr[i] = make([]bool, len(array[i]))
	}
	
	for left <= right && top <= bottom {

		for i := left; i <= right; i++ {
			
			if !traversedArr[left][i] {
				arr = append(arr, array[left][i])
				traversedArr[left][i] = true
			}
			
		}

		for i := top; i < bottom; i++ {
			if !traversedArr[i+1][right] {
				arr = append(arr, array[i+1][right])
				traversedArr[i+1][right] = true
			}
		}

		for i := right; i > left; i-- {
			if !traversedArr[bottom][i-1] {
				arr = append(arr, array[bottom][i-1])
				traversedArr[bottom][i-1] = true
			}
		}

		for i := bottom; i > top+1; i-- {
			if !traversedArr[i-1][left] {
				arr = append(arr, array[i-1][left])
				traversedArr[i-1][left] = true
			}
		}

		left++
		right--
		top++
		bottom--
		
	}
	
	return arr
}
