package main

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
