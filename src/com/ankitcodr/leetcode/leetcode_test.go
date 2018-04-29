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
})
