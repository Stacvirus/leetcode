// Find Closest Number to Zero
// Given an integer array nums of size n, return the number with the value closest to 0 in nums. If there are multiple answers, return the number with the largest value.

// Example 1:

// Input: nums = [-4,-2,1,4,8]
// Output: 1
// Explanation:
// The distance from -4 to 0 is |-4| = 4.
// The distance from -2 to 0 is |-2| = 2.
// The distance from 1 to 0 is |1| = 1.
// The distance from 4 to 0 is |4| = 4.
// The distance from 8 to 0 is |8| = 8.
// Thus, the closest number to 0 in the array is 1.
// Example 2:

// Input: nums = [2,-1,1]
// Output: 1
// Explanation: 1 and -1 are both the closest numbers to 0, so 1 being larger is returned.

package main

import "fmt"

func main() {
	nums := []int{-1000, -1000}
	result := nums[0]

	for i := range nums {
		if abs(nums[i]) < abs(result) {
			result = nums[i]
		}
	}
	closest := 0
	if result < 1 && contains(nums, abs(result)) {
		closest = abs(result)
	} else {
		closest = result
	}
	fmt.Printf("The closest number to zero here is %d\n", closest)
}

func abs(num int) int {
	if num < 1 {
		return num * -1
	}
	return num
}

func contains(slice []int, element int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
