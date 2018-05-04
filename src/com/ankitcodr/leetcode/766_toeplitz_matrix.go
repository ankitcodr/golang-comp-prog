package leetcode

func isToeplitzMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[i])-1; j++ {
			if matrix[i][j] == matrix[i+1][j+1] {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func IsToeplitzMatrix(matrix [][]int) bool {
	return isToeplitzMatrix(matrix)
}
