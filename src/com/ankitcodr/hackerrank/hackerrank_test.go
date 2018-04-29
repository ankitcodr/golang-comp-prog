package hackerrank_test

import (
	"testing"

	. "com/ankitcodr/hackerrank"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLeetcode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Leetcode Suite")
}

var _ = Describe("Leetcode", func() {

	Context("1_Super_Reduced_String", func() {
		It("should pass for default input", func() {
			input := []string{"aaabcc", "aaabccddd", "baab"}
			expected := []string{"ab", "abd", "Empty String"}
			for i:=0;i<len(input);i++ {
				Expect(SuperReducedStringSolution(input[i])).To(Equal(expected[i]))
			}
		})
	})
})
