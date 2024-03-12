package binarygood

import (
	"sort"
	"strings"
)

// largestGood returns the largest good string of S (a good string is a string with an equal number of 0s and 1s)
// The idea is to use a counter to keep track of the number of 1s and 0s, and when the counter reaches 0, we know that we have a good string
// We then recursively call the function on the left and right side of the good string and append the result to the good string
// We then sort the good strings in descending order and join them together
// Time complexity is O(nlogn) because of the sorting
// Space complexity is O(n) because of the recursion stack
func largestGood(S string) string {
	if len(S) == 0 || S[0] == '0' || S[len(S)-1] == '1' {
		return ""
	}
	count := 0
	i := 0
	res := []string{}
	for j := 0; j < len(S); j++ {
		if S[j] == '1' {
			count++
		} else {
			count--
		}
		if count == 0 {
			res = append(res, "1"+largestGood(S[i+1:j])+"0")
			i = j + 1
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] > res[j]
	})
	return strings.Join(res, "")
}
