package leetcode

func TwoSumSolution(nums []int, target int) []int {
	return twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	diffs := make(map[int]int)
	for i, v := range nums {
		start, ok := diffs[target - v]
		if !ok {
			diffs[v] = i
		} else {
			return []int{start, i}
		}
	}
	return []int{}
}