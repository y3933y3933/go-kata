package swap

import "testing"

func TestSwap(t *testing.T) {
	tests := []struct {
		a, b         int
		wantA, wantB int
	}{
		{1, 2, 2, 1},
		{0, 0, 0, 0},
		{100, -50, -50, 100},
	}

	for _, tt := range tests {
		a := tt.a
		b := tt.b
		Swap(&a, &b)
		if a != tt.wantA || b != tt.wantB {
			t.Errorf("Swap(%d, %d) = (%d, %d); want (%d, %d)", tt.a, tt.b, a, b, tt.wantA, tt.wantB)
		}
	}

}
