package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if ind, ok := m[target-nums[i]]; ok {
			return []int{ind, i}

		}
		m[num] = i
	}
	return nil
}
func main() {
	nums := []int{0, 5, 3, 2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	for _, val := range result {
		fmt.Println(val)
	}
}
