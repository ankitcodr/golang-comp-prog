package leetcode_test

import (
	"testing"

	. "com/ankitcodr/leetcode"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLeetcode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Leetcode Suite")
}

var _ = Describe("Leetcode", func() {

	Context("1_Two_Sum", func() {
		It("should pass for positive integers", func() {
			input := []int{2, 7, 11, 15}
			target := 9
			result := []int{0, 1}
			Expect(TwoSumSolution(input, target)).To(Equal(result))
		})
		It("should pass for unsorted integers", func() {
			input := []int{2, 15, 7, 11}
			target := 9
			result := []int{0, 2}
			Expect(TwoSumSolution(input, target)).To(Equal(result))
		})
	})

	Context("20_Valid_Parentheses", func() {
		It("should test stack Push API", func() {
			input := "Ankit"
			output := ""
			st := NewStack()
			for i := 0; i < len(input); i++ {
				st.Push(input[i])
			}
			for st.IsEmpty() != true {
				if val, err := st.Pop(); err == nil {
					temp := string(val) + output
					output = temp
				}
			}
			Expect(output).To(Equal(input))
		})
		It("should pass for default input", func() {
			input := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
			result := []bool{true, true, false, false, true}
			for i := 0; i < len(input); i++ {
				Expect(IsValidSolution(input[i])).To(Equal(result[i]))
			}
		})
	})

	Context("7 Reverse Integer", func() {
		It("should pass for default input", func() {
			input := []int{123, -123, 120}
			output := []int{321, -321, 21}
			for i := 0; i < len(input); i++ {
				Expect(ReverseSolution(input[i])).To(Equal(output[i]))
			}
		})
	})

	Context("26. Remove Duplicates from Sorted Array", func() {
		It("should pass for default input", func() {
			input := [][]int{{}, {1}, {1, 1, 2, 2}, {1, 2, 3}, {0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}
			output := [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {0, 1, 2, 3, 4}}
			for i := 0; i < len(input); i++ {
				Expect(RemoveDuplicatesSolution(input[i])).To(Equal(output[i]))
			}
		})
	})

	Context("27. Remove Element", func() {
		It("should pass for default input", func() {
			input := [][]int{{0}, {0, 0, 1, 2}, {3, 2, 2, 3}, {3, 2, 4, 2, 4, 3}}
			inputVal := []int{0, 0, 3, 4}
			output := [][]int{{}, {1, 2}, {2, 2}, {3, 2, 2, 3}}
			for i := 0; i < len(input); i++ {
				Expect(RemoveElementSolution(input[i], inputVal[i])).To(Equal(output[i]))
			}
		})
	})

	Context("83. Remove Duplicates from Sorted List", func() {
		It("should pass for default input", func() {
			input := ListNode{1, &ListNode{2, &ListNode{3, nil}}}
			output := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
			result := DeleteDuplicatesSolution(&input)
			for output != nil {
				Expect(result.Val).To(Equal(output.Val))
				result = result.Next
				output = output.Next
			}
		})

		It("should remove repeated elements", func() {
			input := ListNode{1, &ListNode{1, &ListNode{3, nil}}}
			output := &ListNode{1, &ListNode{3, nil}}
			result := DeleteDuplicatesSolution(&input)
			for output != nil {
				Expect(result.Val).To(Equal(output.Val))
				result = result.Next
				output = output.Next
			}
		})

		It("should remove repeated elements which are present twice", func() {
			input := ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
			output := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
			result := DeleteDuplicatesSolution(&input)
			for output != nil {
				Expect(result.Val).To(Equal(output.Val))
				result = result.Next
				output = output.Next
			}
		})
	})

	Context("561. Array Partition I", func() {
		It("should pass for default input", func() {
			input := []int{1, 4, 3, 2}
			result := 4
			Expect(ArrayPairSum(input)).To(Equal(result))
		})
	})

	Context("766. Toeplitz Matrix", func() {
		It("should pass for default input", func() {
			input := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
			result := true
			Expect(IsToeplitzMatrix(input)).To(Equal(result))
			input = [][]int{{1, 2}, {2, 2}}
			result = false
			Expect(IsToeplitzMatrix(input)).To(Equal(result))
		})
	})

	Context("566. Reshape the Matrix", func() {
		It("should pass for default input", func() {
			input := [][]int{{1, 2}, {3, 4}}
			result := [][]int{{1, 2, 3, 4}}
			Expect(MatrixReshape(input, 1, 4)).To(Equal(result))
			input = [][]int{{1, 2}, {3, 4}}
			result = [][]int{{1, 2}, {3, 4}}
			Expect(MatrixReshape(input, 2, 4)).To(Equal(result))
		})
	})

	Context("566. Reshape the Matrix", func() {
		It("should pass for default input", func() {
			input := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
			result := 6
			Expect(MaxAreaOfIsland(input)).To(Equal(result))
			input = [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}
			result = 0
			Expect(MaxAreaOfIsland(input)).To(Equal(result))
			input = [][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}
			result = 4
			Expect(MaxAreaOfIsland(input)).To(Equal(result))

		})
	})

	Context("665. Non-decreasing Array", func() {
		It("should pass for default input", func() {
			input := []int{4, 2, 3}
			result := true
			Expect(CheckPossibility(input)).To(Equal(result))
			input = []int{4, 2, 1}
			result = false
			Expect(CheckPossibility(input)).To(Equal(result))
			input = []int{2, 4, 3, 5}
			result = true
			Expect(CheckPossibility(input)).To(Equal(result))
		})
	})
})
