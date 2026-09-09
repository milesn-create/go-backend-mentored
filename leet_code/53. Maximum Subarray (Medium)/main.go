package main

// [1, 6,-4, 8,-5]
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if currentSum+nums[i] > nums[i] {
			currentSum = currentSum + nums[i]
		} else {
			currentSum = nums[i]
		}
		if maxSum < currentSum {
			maxSum = currentSum
		}

	}
	return maxSum

}
