package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxLength := 0

	for i := 0; i < len(s); i++ {
		seen := make(map[byte]bool)
		count := 0

		for j := i; j < len(s); j++ {
			if seen[s[j]] {
				break
			}
			seen[s[j]] = true
			count++
		}

		if count > maxLength {
			maxLength = count
		}
	}

	return maxLength
}
