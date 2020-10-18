package array

// You've given an array of integers and an integer. Write a function that moves all instances of that integer in the array to the end of the array and returns the array

// The function should perform this in place (i.e, it should mutate the input array) and doesn't need to maintain the order of the other intergers

// Sample Input
//
// array = [2, 1, 2, 2, 2, 3, 4, 2]
// toMove = [2]

// Sample Output
//
// [1, 3, 4, 2, 2, 2, 2, 2]

// MoveElementToEnd : array, toMove
func MoveElementToEnd(array []int, toMove int) []int {

	var result []int
	var toMoves []int

	result = make([]int, 0)
	for _, no := range array {
		if no == toMove {
			toMoves = append(toMoves, no)
		} else {
			result = append(result, no)
		}
	}

	result = append(result, toMoves...)

	return result
}
