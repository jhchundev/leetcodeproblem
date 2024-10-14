package main

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 || n == 1 {
		return n
	}

	var lastIndex [256]int
	for i := range lastIndex {
		lastIndex[i] = 0
	}

	maxLen := 0
	start := 0

	for i := 0; i < n; i++ {
		char := s[i]

		if lastIndex[char] > start {
			start = lastIndex[char]
		}

		currentLen := i - start + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}

		lastIndex[char] = i + 1
	}

	return maxLen
}

// test
func main() {
	println(lengthOfLongestSubstring("abcabcbb")) // 3
	println(lengthOfLongestSubstring("bbbbb"))    // 1
	println(lengthOfLongestSubstring("pwwkew"))   // 3
	println(lengthOfLongestSubstring(" "))        // 1
	println(lengthOfLongestSubstring("au"))       // 2
	println(lengthOfLongestSubstring("aab"))      // 2
	println(lengthOfLongestSubstring("dvdf"))     // 3
	println(lengthOfLongestSubstring("abba"))     // 2
	println(lengthOfLongestSubstring("tmmzuxt"))  // 5
	println(lengthOfLongestSubstring("ohvhjdml")) // 6
}
