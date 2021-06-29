package circularbuffer

import (
	"testing"
)

// A circular buffer is a fixed sized buffer into which data can be written. To
// begin with, the buffer is empty. For example, below we have an empty buffer
// with space for five values:
// | . | . | . | . | . |
// As values are written, the buffer fills:
// | 1 | 2 | 3 | . | . |
// and when one more value is added, it looks like:
// | 1 | 2 | 3 | 4 | . |
//   0   1   2   3   4
//   B               E
// When the buffer is full, the next value overwrites the oldest one in the buffer:
// | 6 | 2 | 3 | 4 | 5 |
//       EB
// However, when asked for the contents of the buffer, they are returned in
// insertion order. For the above example, an accessor would return
//[2, 3, 4, 5, 6].

func eq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestCircularBuffer(t *testing.T) {
	in := []int{}
	exp := []int{}
	c := NewCircularBuffer(5)
	got := c.Get()
	if !eq(got, exp) {
		t.Errorf("\nin:%v\nexp:%v\ngot:%v", in, exp, got)
	}

	in = []int{1, 2, 3}
	c.Insert(in)
	exp = []int{1, 2, 3}
	if !eq(c.Get(), exp) {
		t.Errorf("\nin:%v\nexp:%v", in, exp)
	}

	in = []int{4}
	c.Insert(in)
	exp = []int{1, 2, 3, 4}
	if !eq(c.Get(), exp) {
		t.Errorf("\nin:%v\nexp:%v", in, exp)
	}

	in = []int{5, 6}
	c.Insert(in)
	exp = []int{2, 3, 4, 5, 6}
	if !eq(c.Get(), exp) {
		t.Errorf("\nin:%v\nexp:%v", in, exp)
	}
}
