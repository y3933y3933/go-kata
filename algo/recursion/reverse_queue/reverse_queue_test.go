package reversequeue

import (
	"reflect"
	"testing"
)

func TestReverseQueue(t *testing.T) {
	q := []int{1, 2, 3, 4, 5, 6, 7}
	expected := []int{7, 6, 5, 4, 3, 2, 1}

	reverseQueue(&q)

	if !reflect.DeepEqual(q, expected) {
		t.Errorf("Expected %v, got %v", expected, q)
	}
}
