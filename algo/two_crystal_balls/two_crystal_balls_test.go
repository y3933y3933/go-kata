package two_crystal_balls

import "testing"

func TestTwoCrystalBalls(t *testing.T) {
	tests := []struct {
		breaks   []bool
		expected int
	}{
		// 測試資料
		{[]bool{false, false, false, false, true, true, true}, 4},
		{[]bool{false, false, false, false, false, false, false}, -1},
		{[]bool{true, true, true, true, true}, 0},
		{[]bool{false, false, false, true}, 3},
		{[]bool{false, true}, 1},
		{[]bool{true}, 0},
		{[]bool{}, -1},
		{[]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true}, 30},
	}

	for _, tt := range tests {
		result := twoCrystalBalls(tt.breaks)
		if result != tt.expected {
			t.Errorf("TwoCrystalBalls(%v) = %d; want %d", tt.breaks, result, tt.expected)
		}
	}
}
