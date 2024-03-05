package action

import (
	"testing"
)

func TestBottomAction(t *testing.T) {
	testCases := []struct {
		in, want [4][4]int
	}{
		{
			[4][4]int{
				{1, 2, 4, 4},
				{2, 2, 2, 2},
				{3, 4, 2, 4},
				{4, 0, 0, 0},
			},
			[4][4]int{
				{1, 0, 0, 0},
				{2, 0, 0, 4},
				{3, 4, 4, 2},
				{4, 4, 4, 4},
			},
		},
	}
	for _, tc := range testCases {
		result := bottomAction(tc.in)

		if result != tc.want {
			t.Fatalf("Result: %v want %v", result, tc.want)
		}
	}
}
