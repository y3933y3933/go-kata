package counter

import "testing"

func TestCounter(t *testing.T) {
	c := &Counter{}
	c.Increment()
	c.Increment()
	if c.Value() != 2 {
		t.Errorf("Expected 2, got %d", c.Value())
	}
	c.Decrement()
	if c.Value() != 1 {
		t.Errorf("Expected 1, got %d", c.Value())
	}
	c.Reset()
	if c.Value() != 0 {
		t.Errorf("Expected 0, got %d", c.Value())
	}
}
