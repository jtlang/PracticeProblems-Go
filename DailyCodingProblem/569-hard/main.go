package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Problem statement:
    Given an array of numbers of length N, find both the minimum and maximum
    using less than 2 * (N - 2) comparisons.
*/

// The problem breaks down for N < 3
// This solution succeeds where N > 5, and matches the target for N in {4, 5}
// My gut says N in {3, 4, 5} general cases are impossible, but I have not proven this.
func main() {
	var s []int
	if len(os.Args) > 2 {
		for i, a := range os.Args {
			if i > 0 {
				v, _ := strconv.Atoi(a)
				s = append(s, v)
			}
		}
	} else {
		s = []int{1, 2, 3, 4, 5, 6}
	}

	var min, max int

	// One comparison here to initialize min and max with meaningful values
	//  saves two total comparisons against default min/max values later
	count := 1
	if s[0] < s[1] {
		min = s[0]
		max = s[1]
	} else {
		min = s[1]
		max = s[0]
	}

	// Factory captures similar logic - test cond, update old, increment count
	updateIfMin := updateIfBetterFactory(&count, func(old, new int) bool { return new < old })
	updateIfMax := updateIfBetterFactory(&count, func(old, new int) bool { return new > old })

	for i := 2; i+1 < len(s); i += 2 {
		// Critical Section: This trick saves (n-2)/2 comparisons.
		//  Comparing two candidates with each other once enables us to compare
		//  the max with the greater candidate only and the min with the lesser
		//  candidate only, covering this pair in 3 comparisons instead of 4.
		count++
		if s[i] < s[i+1] {
			updateIfMin(&min, s[i])
			updateIfMax(&max, s[i+1])
		} else {
			updateIfMin(&min, s[i+1])
			updateIfMax(&max, s[i])
		}
	}
	if len(s)%2 == 1 {
		// If N is odd, we still need to test the last value
		updateIfMin(&min, s[len(s)-1])
		updateIfMax(&max, s[len(s)-1])
	}

	fmt.Printf("N: %d, Target Comparisons: %d\n", len(s), 2*(len(s)-2))
	fmt.Printf("Min: %d\n", min)
	fmt.Printf("Max: %d\n", max)
	fmt.Printf("Comparisons: %d\n", count)
}

func updateIfBetterFactory(count *int, cond func(int, int) bool) func(*int, int) {
	// Increment comparison count, if condition is satisfied update old with new
	return func(old *int, new int) {
		if cond(*old, new) {
			*old = new
		}
		*count = *count + 1
	}
}
