package array

// Write a function that takes in an array of intergers and returns a boolean representing whether the array is monotonic

// An array is said to be monotonic if its elements, from left to right, are entirely non-increasing or entirely non-decreasing.

// Non-increasing elements aren't necessarily exclusively decreasing; they simply don't increase. Similarly, non-decreasing elements
// aren't necessarily exclusively increasing; they simply don't decrease.

// Note that empty arrays and arrays of one element are monotonic.

// Sample Input
//
// array = [-1, -5, -10, -1100, -1100, -1101, -1102, -9001]
//
// Sample Output
//
// true

// TrendingType : upward / downward
type TrendingType int

const (
	TrendingTypeNone TrendingType = iota
	TrendingTypeUpward 
	TrendingTypeDownward
)

// IsMonotonic : array
func IsMonotonic(array []int) bool {
	if len(array) == 0 || len(array) == 1 {
		return true
	}

	var value int
	var trend TrendingType
	for i, no := range array {

		if i == 0 {
			value = no
			continue
		} 
		if (no > value) {
			if trend == TrendingTypeNone {
				trend = TrendingTypeUpward
			} else if(trend == TrendingTypeDownward) {
				return false
			}
		} else if(no < value) {
			if trend == TrendingTypeNone {
				trend = TrendingTypeDownward
			} else if(trend == TrendingTypeUpward) {
				return false
			}
		}
		
		value = no
	}
	
	return true
}

