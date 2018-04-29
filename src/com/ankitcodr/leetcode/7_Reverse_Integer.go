package leetcode

const (
	INT_MAX = 2 * 1024 * 1024 * 1024
)

func reverse(x int) int {
	copy := int64(x)
	if x < 0 {
		copy = copy * -1
	}
	result := int64(0)
	for copy != 0 {
		result = (result * 10) + copy%10
		copy /= 10
	}
	if result > INT_MAX {
		return 0
	}
	if x < 0 {
		return -1 * int(result)
	}
	return int(result)
}

func ReverseSolution(x int) int {
	return reverse(x)
}
