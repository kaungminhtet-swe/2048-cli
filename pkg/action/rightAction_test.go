package action

import (
	"testing"
)

func TestRightAction(t *testing.T) {
	testcases := []struct {
		in, want [4][4]int
	}{
		{
			[4][4]int{
				{1, 2, 3, 4},
				{2, 2, 4, 0},
				{4, 2, 2, 0},
				{4, 2, 4, 0},
			},
			[4][4]int{
				{1, 2, 3, 4},
				{0, 0, 4, 4},
				{0, 0, 4, 4},
				{0, 4, 2, 4},
			},
		},
	}
	for _, tc := range testcases {
		result := rightAction(tc.in)
		if result != tc.want {
			t.Fatalf("Result: %v actual: %v", result, tc.want)
		}
	}
}
