package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {

	testCases := []TestCase{
		{
			Name: "testcase",
			Data: "babad",
			Want: "bab",
		},
		{
			Name: "testcase",
			Data: "cbbd",
			Want: "bb",
		},
	}
	for i, tc := range testCases {
		t.Run(tc.Name+string(rune(i)), func(t *testing.T) {
			got := LongestPalindrome(tc.Data)
			assert.Equal(t, tc.Want, got)
		})
	}
}
