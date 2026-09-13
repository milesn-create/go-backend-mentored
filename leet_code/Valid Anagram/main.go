package main

import "sort"

func isAnagram(s string, t string) bool {
	m := make(map[rune]int)
	for _, cimbol1 := range s {
		if count, ok := m[cimbol1]; ok {
			count++
			m[cimbol1] = count
		} else {
			m[cimbol1] = 1
		}
	}
	for _, cimbol2 := range t {
		if count, ok := m[cimbol2]; ok {
			count--
			m[cimbol2] = count
		} else {
			return false
		}
	}
	for _, count := range m {
		if count != 0 {
			return false

		}

	}
	return true
}
func IsAnagram(s string, t string) bool {
	runes1 := []rune(s)
	sort.Slice(runes1, func(i, j int) bool {
		return runes1[i] < runes1[j]
	})
	runes2 := []rune(t)
	sort.Slice(runes2, func(i, j int) bool {
		return runes2[i] < runes2[j]
	})
	if string(runes1) == string(runes2) {
		return true
	}
	return false

}
