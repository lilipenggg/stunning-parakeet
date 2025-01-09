package main

import "fmt"

func isPrefixAndSuffix(str1, str2 string) bool {
	// 'aba', 'ababa'
	if str1 == str2[0:len(str1)] && str1 == str2[len(str2)-len(str1):] {
		return true
	}

	return false
}

func countPrefixSuffixPairs(words []string) int {
	count := 0

	for i, w := range words {
		// not the last word in the slice
		if i != len(words)-1 {
			remainingWords := words[i+1:]

			for j, rw := range remainingWords {
				if len(w) <= len(rw) && isPrefixAndSuffix(w, rw) {
					count++
					fmt.Printf("i = %d and j = %d\n", i, i+j+1)
				}
			}
		}
	}

	fmt.Printf("count: %d\n", count)

	return count
}

func main() {
	countPrefixSuffixPairs([]string{"a", "aba", "ababa", "aa"})
	countPrefixSuffixPairs([]string{"pa", "papa", "ma", "mama"})
	countPrefixSuffixPairs([]string{"abab", "ab"})
}
