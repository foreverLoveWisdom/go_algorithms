package main

import "testing"

func TestLongestSubstringWithKDistinctExample1(t *testing.T) {
	s := "abcba"
	k := 2
	want := 3
	got := lengthOfLongestSubstringKDistinct(s, k)
	if got != want {
		t.Errorf("Test failed for input %v with k=%d: expected %v, got %v", s, k, want, got)
	}
}

func TestLongestSubstringWithKDistinctExample2(t *testing.T) {
	s := "eceba"
	k := 2
	want := 3 // Longest substring is "ece"
	got := lengthOfLongestSubstringKDistinct(s, k)
	if got != want {
		t.Errorf("Test failed for input %v with k=%d: expected %v, got %v", s, k, want, got)
	}
}

func TestLongestSubstringWithKDistinctEdgeCaseKZero(t *testing.T) {
	s := "abcba"
	k := 0
	want := 0 // No valid substrings when k = 0
	got := lengthOfLongestSubstringKDistinct(s, k)
	if got != want {
		t.Errorf("Test failed for input %v with k=%d: expected %v, got %v", s, k, want, got)
	}
}

func TestLongestSubstringWithKDistinctEdgeCaseEmptyString(t *testing.T) {
	s := ""
	k := 2
	want := 0 // No substrings in an empty string
	got := lengthOfLongestSubstringKDistinct(s, k)
	if got != want {
		t.Errorf("Test failed for input %v with k=%d: expected %v, got %v", s, k, want, got)
	}
}

func TestLongestSubstringWithKDistinctEdgeCaseKLargerThanString(t *testing.T) {
	s := "abc"
	k := 5
	want := 3 // Whole string is valid since k > len(s)
	got := lengthOfLongestSubstringKDistinct(s, k)
	if got != want {
		t.Errorf("Test failed for input %v with k=%d: expected %v, got %v", s, k, want, got)
	}
}

func TestLongestSubstringWithKDistinctSingleChar(t *testing.T) {
	s := "aaaaa"
	k := 1
	want := 5 // Whole string is valid
	got := lengthOfLongestSubstringKDistinct(s, k)
	if got != want {
		t.Errorf("Test failed for input %v with k=%d: expected %v, got %v", s, k, want, got)
	}
}

func TestLongestSubstringWithKDistinctAllDistinct(t *testing.T) {
	s := "abcdef"
	k := 3
	want := 3 // Longest substring with at most 3 distinct characters is "abc"
	got := lengthOfLongestSubstringKDistinct(s, k)
	if got != want {
		t.Errorf("Test failed for input %v with k=%d: expected %v, got %v", s, k, want, got)
	}
}

func TestLongestSubstringWithKDistinctKIsOne(t *testing.T) {
	s := "abaccc"
	k := 1
	want := 3 // Longest substring is "ccc"
	got := lengthOfLongestSubstringKDistinct(s, k)
	if got != want {
		t.Errorf("Test failed for input %v with k=%d: expected %v, got %v", s, k, want, got)
	}
}
