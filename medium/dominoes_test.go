package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushDominoes(t *testing.T) {
	testCases := []TestCase{
		{
			Name: "testcase",
			Data: "RR.L",
			Want: "RR.L",
		},
		{
			Name: "testcase",
			Data: ".L.R...LR..L..",
			Want: "LL.RR.LLRRLL..",
		},
		{
			Name: "testcase",
			Data: ".L.R.",
			Want: "LL.RR",
		},
		{
			Name: "testcase",
			Data: ".",
			Want: ".",
		},
		{
			Name: "testcase",
			Data: "L",
			Want: "L",
		},
		{
			Name: "testcase",
			Data: "R",
			Want: "R",
		},
		{
			Name: "runs",
			Data: "R.",
			Want: "RR",
		},
		{
			Name: "testcase",
			Data: "..R..",
			Want: "..RRR",
		},
	}
	for i, tc := range testCases {
		t.Run(tc.Name+string(rune(i)), func(t *testing.T) {
			got := PushDominoes(tc.Data)
			assert.Equal(t, tc.Want, got)
		})
	}
}
