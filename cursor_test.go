package codegraph

import "testing"

var input = []int{
	1, 11, 21, 31, 41, 51, 61, 71, 81, 91, 101,
}

func TestCursor(t *testing.T) {
	var c Cursor[int]
	t.Run("New", func(t *testing.T) {
		c = NewCursor(input)
		if c == nil {
			t.Errorf("expected cursor not to be nil")
		}
		t.Run("Fail", func(t *testing.T) {
			n := NewCursor([]int{})
			if n != nil {
				t.Errorf("expected cursor to be nil")
			}
		})
	})
	t.Run("Cur", func(t *testing.T) {
		if c.Cur() != input[0] {
			t.Errorf("unexpected value: wanted %d ; got %d", input[0], c.Cur())
		}
	})
	t.Run("Next", func(t *testing.T) {
		v := c.Next()
		if v != input[1] {
			t.Errorf("unexpected value: wanted %d ; got %d", input[1], v)
		}
	})
	t.Run("Prev", func(t *testing.T) {
		v := c.Prev()
		if v != input[0] {
			t.Errorf("unexpected value: wanted %d ; got %d", input[0], v)
		}
		t.Run("HitHead", func(t *testing.T) {
			v := c.Prev()
			if v != input[0] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[0], v)
			}
		})
		// advance to idx 1
		c.Next()
	})
	t.Run("Peek", func(t *testing.T) {
		v := c.Peek()
		w := c.PeekOffset(1)
		x := c.PeekOffset(0)
		cur := c.Cur()

		if v != input[2] {
			t.Errorf("unexpected value: wanted %d ; got %d", input[2], v)
		}
		if w != v {
			t.Error("ambiguous")
		}
		if x != cur {
			t.Error("ambiguous")
		}
		if cur != input[1] {
			t.Errorf("unexpected value: wanted %d ; got %d", input[1], v)
		}
		t.Run("HitTail", func(t *testing.T) {
			c.Idx(10)
			v := c.Peek()
			cur := c.Cur()
			if v != input[10] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[10], v)
			}
			if cur != input[10] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[10], v)
			}
		})
		c.Idx(2)
	})
	t.Run("PeekIdx", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			v := c.PeekIdx(4)
			cur := c.Cur()
			if v != input[4] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[4], v)
			}
			if cur != input[2] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[2], v)
			}
		})
		t.Run("HitHead", func(t *testing.T) {
			v := c.PeekIdx(-10)
			cur := c.Cur()
			if v != input[0] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[0], v)
			}
			if cur != input[2] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[2], v)
			}
		})
		t.Run("HitTail", func(t *testing.T) {
			v := c.PeekIdx(20)
			cur := c.Cur()
			if v != input[10] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[10], v)
			}
			if cur != input[2] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[2], v)
			}
		})
	})

	t.Run("PeekOffset", func(t *testing.T) {
		t.Run("SuccessPositive", func(t *testing.T) {
			v := c.PeekOffset(2)
			cur := c.Cur()
			if v != input[4] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[4], v)
			}
			if cur != input[2] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[2], v)
			}
		})
		t.Run("SuccessNegative", func(t *testing.T) {
			v := c.PeekOffset(-1)
			cur := c.Cur()
			if v != input[1] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[1], v)
			}
			if cur != input[2] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[2], v)
			}
		})
		t.Run("HitHead", func(t *testing.T) {
			v := c.PeekOffset(-10)
			cur := c.Cur()
			if v != input[0] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[0], v)
			}
			if cur != input[2] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[2], v)
			}
		})
		t.Run("HitTail", func(t *testing.T) {
			v := c.PeekOffset(20)
			cur := c.Cur()
			if v != input[10] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[10], v)
			}
			if cur != input[2] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[2], v)
			}
		})
	})
	t.Run("Pos", func(t *testing.T) {
		p := c.Pos()
		if p != 2 {
			t.Errorf("unexpected position: wanted %d ; got %d", 2, p)
		}
	})
	t.Run("Pos", func(t *testing.T) {
		l := c.Len()
		if l != 11 {
			t.Errorf("unexpected position: wanted %d ; got %d", 11, l)
		}
	})
	t.Run("Head", func(t *testing.T) {
		v := c.Head()
		if v != input[0] {
			t.Errorf("unexpected value: wanted %d ; got %d", input[0], v)
		}
	})
	t.Run("Tail", func(t *testing.T) {
		v := c.Tail()
		if v != input[10] {
			t.Errorf("unexpected value: wanted %d ; got %d", input[10], v)
		}
	})
	t.Run("Idx", func(t *testing.T) {
		v := c.Idx(5)
		if v != input[5] {
			t.Errorf("unexpected value: wanted %d ; got %d", input[5], v)
		}

		t.Run("HitTail", func(t *testing.T) {
			v := c.Idx(20)
			if v != input[10] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[10], v)
			}
		})
		t.Run("HitHead", func(t *testing.T) {
			v := c.Idx(-1)
			if v != input[0] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[0], v)
			}
		})
		c.Idx(5)
	})
	t.Run("Offset", func(t *testing.T) {
		t.Run("SuccessPositive", func(t *testing.T) {
			v := c.Offset(2)
			if v != input[7] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[7], v)
			}
		})
		t.Run("SuccessNegative", func(t *testing.T) {
			v := c.Offset(-3)
			if v != input[4] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[4], v)
			}
		})

		t.Run("HitTail", func(t *testing.T) {
			v := c.Offset(8)
			if v != input[10] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[10], v)
			}
		})
		t.Run("HitHead", func(t *testing.T) {
			v := c.Offset(-30)
			if v != input[0] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[0], v)
			}
		})
	})
}
