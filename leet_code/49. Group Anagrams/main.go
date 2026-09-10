package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		key := string(runes)
		groups[key] = append(groups[key], str)
	}
	result := [][]string{}
	for _, group := range groups {
		result = append(result, group)
	}
	return result

}
