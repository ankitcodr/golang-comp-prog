package leetcode

func checkPossibility(nums []int) bool {
	if len(nums) < 3 {
		return true
	}
	cnt := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if cnt >= 1 {
				return false
			}
			cnt++
			if i==1 {
				nums[i-1] = nums[i]
				continue
			}
			if nums[i] >= nums[i-2] {
				nums[i-1] = nums[i]
			} else {
				nums[i]= nums[i-1]
			}
		}
	}
	return true
}

func CheckPossibility(nums []int) bool {
	return checkPossibility(nums)
}
