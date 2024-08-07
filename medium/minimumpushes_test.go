package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinumumPushes(t *testing.T) {
	testCases := []TestCase{
		{
			Name: "testcase 1",
			Data: "abcde",
			Want: 5,
		},
		{
			Name: "testcase 2",
			Data: "xyzxyzxyzxyz",
			Want: 12,
		},
		{
			Name: "testcase 3",
			Data: "aabbccddeeffgghhiiiiii",
			Want: 24,
		},
		{
			Name: "testcase 5",
			Data: "au",
			Want: 2,
		},
		{
			Name: "testcase 6",
			Data: "abcb",
			Want: 4,
		},
		{
			Name: "testcase 7",
			Data: "afhtgpque",
			Want: 10,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			got := MinimumPushes(tc.Data)
			assert.Equal(t, tc.Want, got)
		})
	}
}
