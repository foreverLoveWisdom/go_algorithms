package main

import "testing"

// Given a string s, find the length of the longest substring
// without repeating characters.

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

// Constraints:

//     0 <= s.length <= 5 * 104
//     s consists of English letters, digits, symbols and spaces.

func TestLongestSubstringExample0(t *testing.T) {
	s := "abba"
	want := 2
	got := lengthOfLongestSubstring(s)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", s, want, got)
	}
}

func TestLongestSubstringExample1(t *testing.T) {
	s := "abac"
	want := 3
	got := lengthOfLongestSubstring(s)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", s, want, got)
	}
}

func TestLongestSubstringExample2(t *testing.T) {
	s := "bbbbb"
	want := 1
	got := lengthOfLongestSubstring(s)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", s, want, got)
	}
}

func TestLongestSubstringExample3(t *testing.T) {
	s := "pwwkew"
	want := 3
	got := lengthOfLongestSubstring(s)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", s, want, got)
	}
}

func TestLongestSubstringEmptyString(t *testing.T) {
	s := ""
	want := 0
	got := lengthOfLongestSubstring(s)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", s, want, got)
	}
}

func TestLongestSubstringSingleCharacter(t *testing.T) {
	s := "a"
	want := 1
	got := lengthOfLongestSubstring(s)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", s, want, got)
	}
}

func TestLongestSubstringAllUnique(t *testing.T) {
	s := "ab"
	want := 2
	got := lengthOfLongestSubstring(s)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", s, want, got)
	}
}

func TestLongestSubstringWithSymbols(t *testing.T) {
	s := "a1!@"
	want := 4
	got := lengthOfLongestSubstring(s)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", s, want, got)
	}
}
