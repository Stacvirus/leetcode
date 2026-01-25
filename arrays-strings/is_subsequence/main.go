// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

// Example 1:

// Input: s = "abc", t = "ahbgdc"
// Output: true
// Example 2:

// Input: s = "axc", t = "ahbgdc"
// Output: false

// Constraints:

// 0 <= s.length <= 100
// 0 <= t.length <= 104
// s and t consist only of lowercase English letters.

package main

import "strings"

func main() {
	s := "axc"
	t := "ahbgdc"
	i, j := 0, 0
	var found strings.Builder

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			found.WriteByte(s[i])
			i++
		}
		j++
	}
	if found.Len() == len(s) {
		println("true")
		return
	}
	println("false")

}
