package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	str := "abcabcbb"
	res := LengthOfLongestSubstring(str)
	assert.Equal(t, 3, res)

	str = "bbbbb"
	res = LengthOfLongestSubstring(str)
	assert.Equal(t, 1, res)

	str = "pwwkew"
	res = LengthOfLongestSubstring(str)
	assert.Equal(t, 3, res)

	str = "cdd"
	res = LengthOfLongestSubstring(str)
	assert.Equal(t, 2, res)

	str = "abba"
	res = LengthOfLongestSubstring(str)
	assert.Equal(t, 2, res)

	str = "abcb"
	res = LengthOfLongestSubstring(str)
	assert.Equal(t, 3, res)
}
