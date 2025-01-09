package main

import (
	"fmt"
	"math"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	if len(nums) == 0 || (len(nums) == 1 && nums[0] == 0) {
		return 0
	}

	median := float64(0)

	// sort nums
	sort.Ints(nums)

	if math.Mod(float64(len(nums)), 2) == 0 {
		// even length, e.g. [1,2,3,4]
		rightIdx := len(nums) / 2
		leftIdx := rightIdx - 1

		median = float64(nums[leftIdx]+nums[rightIdx]) / float64(2)

	} else {
		// odd length, e.g. [1,2,3], [1,2,3,4,5], [1,2,3,4,5,6,7]
		idx := int(math.Floor(float64(len(nums)) / float64(2)))

		median = float64(nums[idx])
	}

	fmt.Println("median:", median)

	return median
}

func main() {
	findMedianSortedArrays([]int{1, 3}, []int{2})
	findMedianSortedArrays([]int{1, 2}, []int{3, 4})
}
