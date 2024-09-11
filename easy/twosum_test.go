package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name  string
	Data  []int
	Data2 int
	Want  []int
}

func TestTwoSum(t *testing.T) {
	testCases := []TestCase{
		{
			Name:  "testcase",
			Data:  []int{0, 4, 3, 0}, // ..R.LL
			Data2: 0,
			Want:  []int{0, 3},
		},
		{
			Name:  "testcase",
			Data:  []int{3, 2, 4}, // ..R.LL
			Data2: 6,
			Want:  []int{1, 2},
		},
		{
			Name:  "testcase",
			Data:  []int{2, 7, 11, 15}, // ..R.LL
			Data2: 9,
			Want:  []int{0, 1},
		},
		{
			Name:  "testcase",
			Data:  []int{3, 3}, // ..R.LL
			Data2: 6,
			Want:  []int{0, 1},
		},
	}
	for i, tc := range testCases {
		t.Run(tc.Name+string(rune(i)), func(t *testing.T) {
			got := TwoSum(tc.Data, tc.Data2)
			assert.Equal(t, tc.Want, got)
		})
	}
}
