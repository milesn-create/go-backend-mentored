package main

import "unicode"

func isPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(s)-1; i < j; {
		if unicode.IsDigit(runes[i]) || unicode.IsLetter(runes[i]) {
			if unicode.IsDigit(runes[j]) || unicode.IsLetter(runes[j]) {
				if unicode.ToLower(runes[i]) == unicode.ToLower(runes[j]) {
					i++
					j--
				} else {
					return false
				}
			} else {
				j--
			}
		} else {
			i++
		}
	}
	return true
}
