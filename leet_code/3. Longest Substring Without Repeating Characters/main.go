package main

func lengthOfLongestSubstring(s string) int {
	left := 0
	maxLenSubstring := 0
	m := make(map[rune]int)
	str := []rune(s)
	for right := 0; right < len(str); right++ {
		simbol := str[right]
		if ind, ok := m[simbol]; ok && ind >= left {
			left = ind + 1

		}
		m[simbol] = right
		if right-left+1 > maxLenSubstring {
			maxLenSubstring = right - left + 1
		}

	}
	return maxLenSubstring

}
