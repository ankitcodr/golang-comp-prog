package leetcode

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums)*len(nums[0]) != r*c {
		return nums
	}

	var result [][]int
	result = make([][]int, r)
	result[0] = make([]int, c)
	var k, l int

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if l >= c {
				k++
				result[k] = make([]int, c)
				l = 0
			}
			result[k][l] = nums[i][j]
			l++
		}
	}
	return result
}

func MatrixReshape(nums [][]int, r int, c int) [][]int {
	return matrixReshape(nums, r, c)
}
