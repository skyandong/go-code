package leetcode

import "fmt"

func lengthOfLongestSubstring(s string) int {
	mymap := make(map[uint8]int)
	var max int
	var first int
	for end := 1;end < len(s);end++{
		mymap[s[end]]++
		if first == end{
			continue
		}

	}
}
