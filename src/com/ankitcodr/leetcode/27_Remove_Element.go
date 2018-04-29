package leetcode

func removeElement(nums []int, val int) int {
	l := len(nums)
	if l <= 0 {
		return l
	}
	i := 0
	for j := i; j < l; j++ {
		if val == nums[j] {
			continue
		} else {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func RemoveElementSolution(nums []int, val int) []int {
	l := removeElement(nums, val)
	nums = nums[:l]
	return nums
}
