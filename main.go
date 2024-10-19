package main

import (
	"math"
)

func MaximumRescue(roof int, chicksPositions []int) int {
	max := maxItemInRange(chicksPositions, roof, 0, len(chicksPositions)-1)
	return max
}

// A function used to looking for maximum item that in range of pRange
// The concept is using divide-and-conquer algorithm to split the positions in halves
// And process that splited array in each like a binary search
// Time Complexity is O(n log n) because O(log n) for divide-and-conquer algorithm
// And O(n) for each side (left and right) iteration to find max value
func maxItemInRange(positions []int, pRange int, start, end int) int {
	// Prevent the last position case
	if start >= len(positions) {
		return 0
	}
	// Picked the mid position for search and split
	mid := midPosition(start, end)

	// Find in left side
	leftCounts := maxLeftSide(positions, pRange, mid)
	// Find in right side
	rightCounts := maxRightSide(positions, pRange, mid)

	maxCount := max(leftCounts, rightCounts)

	// Split left and right to process further
	if end > start {
		leftCounts = maxItemInRange(positions, pRange, start, mid-1)
		rightCounts = maxItemInRange(positions, pRange, mid+1, end)
		maxCount = max(maxCount, max(leftCounts, rightCounts))
	}

	return maxCount
}

// Helper max function
func max(a, b int) int {
	max := a
	if b > a {
		max = b
	}
	return max
}

func midPosition(start, end int) int {
	return int(math.Ceil(float64(start) + (float64(end)-float64(start))/2))
}

// The logic for left side
// Time Complexity: O(n)
func maxLeftSide(arr []int, pRange, mid int) int {
	leftCounts := 0
	// the limit checkpoint: limit <- mid
	leftToMid := (arr[mid] + 1) - pRange

	for i := mid; i >= 0; i-- {
		// break when exceeded limit
		if arr[i] < leftToMid {
			break
		}

		leftCounts += 1
	}
	return leftCounts
}

// The logic for right side
// Time Complexity: O(n)
func maxRightSide(arr []int, pRange, mid int) int {
	rightCounts := 0
	// the limit checkpoint: mid -> limit
	midToRight := (arr[mid] - 1) + pRange

	for i := mid; i < len(arr); i++ {
		// break when exceeded limit
		if arr[i] > midToRight {
			break
		}

		rightCounts += 1
	}
	return rightCounts
}
