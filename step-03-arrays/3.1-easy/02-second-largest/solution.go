package secondlargest

import "math"

// SecondLargest returns the second largest distinct value, or -1 if there is none.
// One pass. Time O(n), space O(1).
func SecondLargest(a []int) int {
	largest, second := math.MinInt, math.MinInt
	for _, x := range a {
		if x > largest {
			second = largest
			largest = x
		} else if x < largest && x > second {
			second = x
		}
	}
	if second == math.MinInt {
		return -1
	}
	return second
}

// SecondSmallest returns the second smallest distinct value, or -1 if there is none.
func SecondSmallest(a []int) int {
	smallest, second := math.MaxInt, math.MaxInt
	for _, x := range a {
		if x < smallest {
			second = smallest
			smallest = x
		} else if x > smallest && x < second {
			second = x
		}
	}
	if second == math.MaxInt {
		return -1
	}
	return second
}
