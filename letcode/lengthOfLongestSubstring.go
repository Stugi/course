package main

import (
	"fmt"
	"strings"
)

// длина самой длинной подстроки
//Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanati on: The answer is "abc", with the length of 3.
func main() {
	fmt.Println(lengthOfLongestSubstring("asjrgapa"))
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	var tmpMap []string
	var indMaxElement = 0
	tmpMap = make([]string, 0, 10)
	for ind, el := range s {
		tmpMap = append(tmpMap, string(el))
		for ind1, _ := range tmpMap {
			if ind != ind1 {
				if (ind <= len(tmpMap[ind1])+ind1) && !strings.Contains(tmpMap[ind1], string(el)) {
					tmpMap[ind1] = tmpMap[ind1] + string(el)
				}
				if len(tmpMap[indMaxElement]) <= len(tmpMap[ind1]) {
					indMaxElement = ind1
				}
			}
		}
	}

	return len(tmpMap[indMaxElement])
}
