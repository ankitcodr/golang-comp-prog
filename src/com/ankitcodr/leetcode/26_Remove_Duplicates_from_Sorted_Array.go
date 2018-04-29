package leetcode

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}
	i := 0
	for j := i + 1; j < l; j++ {
		if nums[i] == nums[j] {
			continue
		} else {
			nums[i+1] = nums[j]
			i++
		}
	}
	return i + 1
}

func RemoveDuplicatesSolution(nums []int) []int {
	l := removeDuplicates(nums)
	nums = nums[:l]
	return nums
}
