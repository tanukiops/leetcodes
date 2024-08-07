package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name string
	Data string
	Want any
}

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []TestCase{
		{
			Name: "testcase 1",
			Data: "abcabcbb",
			Want: 3,
		},
		{
			Name: "testcase 2",
			Data: "bbbbb",
			Want: 1,
		},
		{
			Name: "testcase 3",
			Data: "pwwkew",
			Want: 3,
		},
		{
			Name: "testcase 4",
			Data: " ",
			Want: 1,
		},
		{
			Name: "testcase 5",
			Data: "au",
			Want: 2,
		},
		{
			Name: "testcase 6",
			Data: "",
			Want: 0,
		},
		{
			Name: "testcase 7",
			Data: "abcb",
			Want: 3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			got := LenghtOfLongestSubstring(tc.Data)
			assert.Equal(t, tc.Want, got)
		})
	}
}
