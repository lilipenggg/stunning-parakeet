package main

import "fmt"

func twoSumBruteForce(nums []int, target int) []int {
	// O(n^2)
	for i, rn := range nums {
		// not the last element
		if i != len(nums)-1 {
			remainNums := nums[i+1:]
			for j, ln := range remainNums {
				if rn+ln == target {
					return []int{i, i + 1 + j}
				}
			}
		}
	}
	return nil
}

func twoSumHashMap(nums []int, target int) []int {
	// O(n)
	hashMap := map[int]int{}
	for i, n := range nums {
		if idx, ok := hashMap[target-n]; ok {
			return []int{idx, i}
		} else {
			hashMap[n] = i
		}
	}
	return nil
}

func main() {
	result := twoSumBruteForce([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)
	result = twoSumBruteForce([]int{3, 2, 4}, 6)
	fmt.Println(result)
	result = twoSumBruteForce([]int{3, 3}, 6)
	fmt.Println(result)

	result = twoSumHashMap([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)
	result = twoSumHashMap([]int{3, 2, 4}, 6)
	fmt.Println(result)
	result = twoSumHashMap([]int{3, 3}, 6)
	fmt.Println(result)
}
