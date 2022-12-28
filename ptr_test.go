package cur

import "testing"

var (
	ptrinput = &[]int{}
	newData  = []int{111, 121, 131}
)

func unset() {
	*ptrinput = []int{}
}
func set() {
	*ptrinput = []int{
		1, 11, 21, 31, 41, 51, 61, 71, 81, 91, 101,
	}
}
func add() {
	*ptrinput = append(*ptrinput, newData...)
}

func TestPointer(t *testing.T) {
	var c Cursor[int]
	set()

	t.Run("Ptr", func(t *testing.T) {
		c = Ptr(ptrinput)
		if c == nil {
			t.Errorf("expected cursor not to be nil")
		}
		t.Run("Initialized", func(t *testing.T) {
			var initialized = &[]int{}
			n := Ptr(initialized)
			if n == nil {
				t.Errorf("expected cursor not to be nil")
			}
		})
		t.Run("Fail", func(t *testing.T) {
			var zero *[]int
			n := Ptr(zero)
			if n != nil {
				t.Errorf("expected cursor to be nil")
			}
		})
	})

	t.Run("Cur", func(t *testing.T) {
		if c.Cur() != input[0] {
			t.Errorf("unexpected value: wanted %d ; got %d", input[0], c.Cur())
		}

		t.Run("Fail", func(t *testing.T) {
			unset()
			if c.Cur() != 0 {
				t.Errorf("unexpected value: wanted %d ; got %d", 0, c.Cur())
			}
			set()
		})
	})

	t.Run("Next", func(t *testing.T) {
		v := c.Next()
		if v != input[1] {
			t.Errorf("unexpected value: wanted %d ; got %d", input[1], v)
		}
		t.Run("Fail", func(t *testing.T) {
			unset()
			if c.Next() != 0 {
				t.Errorf("unexpected value: wanted %d ; got %d", 0, c.Next())
			}
			set()
		})
	})
	t.Run("Prev", func(t *testing.T) {
		v := c.Prev()
		if v != input[0] {
			t.Errorf("unexpected value: wanted %d ; got %d", input[0], v)
		}
		t.Run("HitHead", func(t *testing.T) {
			v := c.Prev()
			if v != eof {
				t.Errorf("unexpected value: wanted %d ; got %d", eof, v)
			}
		})
		t.Run("Fail", func(t *testing.T) {
			unset()
			if c.Prev() != 0 {
				t.Errorf("unexpected value: wanted %d ; got %d", 0, c.Prev())
			}
			set()
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
			if v != eof {
				t.Errorf("unexpected value: wanted %d ; got %d", input[10], v)
			}
			if cur != input[10] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[10], v)
			}
		})
		t.Run("Fail", func(t *testing.T) {
			unset()
			if c.Peek() != 0 {
				t.Errorf("unexpected value: wanted %d ; got %d", 0, c.Peek())
			}
			set()
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
			if v != eof {
				t.Errorf("unexpected value: wanted %d ; got %d", input[0], v)
			}
			if cur != input[2] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[2], v)
			}
		})
		t.Run("HitTail", func(t *testing.T) {
			v := c.PeekIdx(20)
			cur := c.Cur()
			if v != eof {
				t.Errorf("unexpected value: wanted %d ; got %d", input[10], v)
			}
			if cur != input[2] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[2], v)
			}
		})
		t.Run("Fail", func(t *testing.T) {
			unset()
			if c.PeekIdx(3) != 0 {
				t.Errorf("unexpected value: wanted %d ; got %d", 0, c.PeekIdx(3))
			}
			set()
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
			if v != eof {
				t.Errorf("unexpected value: wanted %d ; got %d", input[0], v)
			}
			if cur != input[2] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[2], v)
			}
		})
		t.Run("HitTail", func(t *testing.T) {
			v := c.PeekOffset(20)
			cur := c.Cur()
			if v != eof {
				t.Errorf("unexpected value: wanted %d ; got %d", input[10], v)
			}
			if cur != input[2] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[2], v)
			}
		})
		t.Run("Fail", func(t *testing.T) {
			unset()
			if c.PeekOffset(3) != 0 {
				t.Errorf("unexpected value: wanted %d ; got %d", 0, c.PeekOffset(3))
			}
			set()
		})
	})
	t.Run("Pos", func(t *testing.T) {
		p := c.Pos()
		if p != 2 {
			t.Errorf("unexpected position: wanted %d ; got %d", 2, p)
		}
		t.Run("Fail", func(t *testing.T) {
			unset()
			if c.Pos() != -1 {
				t.Errorf("unexpected value: wanted %d ; got %d", -1, c.Pos())
			}
			set()
		})
	})
	t.Run("Len", func(t *testing.T) {
		l := c.Len()
		if l != 11 {
			t.Errorf("unexpected position: wanted %d ; got %d", 11, l)
		}
		t.Run("Fail", func(t *testing.T) {
			unset()
			if c.Len() != 0 {
				t.Errorf("unexpected value: wanted %d ; got %d", 0, c.Len())
			}
			set()
		})
	})
	t.Run("Head", func(t *testing.T) {
		v := c.Head()
		if v != input[0] {
			t.Errorf("unexpected value: wanted %d ; got %d", input[0], v)
		}
		t.Run("Fail", func(t *testing.T) {
			unset()
			if c.Head() != 0 {
				t.Errorf("unexpected value: wanted %d ; got %d", 0, c.Head())
			}
			set()
		})
	})
	t.Run("Tail", func(t *testing.T) {
		v := c.Tail()
		if v != input[10] {
			t.Errorf("unexpected value: wanted %d ; got %d", input[10], v)
		}
		t.Run("Fail", func(t *testing.T) {
			unset()
			if c.Tail() != 0 {
				t.Errorf("unexpected value: wanted %d ; got %d", 0, c.Tail())
			}
			set()
		})
		t.Run("UpdatedPointer", func(t *testing.T) {
			add()
			if c.Tail() != 131 {
				t.Errorf("unexpected value: wanted %d ; got %d", 131, c.Tail())
			}
			set()
		})
	})
	t.Run("Idx", func(t *testing.T) {
		v := c.Idx(5)
		if v != input[5] {
			t.Errorf("unexpected value: wanted %d ; got %d", input[5], v)
		}

		t.Run("HitTail", func(t *testing.T) {
			v := c.Idx(20)
			if v != eof {
				t.Errorf("unexpected value: wanted %d ; got %d", input[10], v)
			}
		})
		t.Run("HitHead", func(t *testing.T) {
			v := c.Idx(-1)
			if v != eof {
				t.Errorf("unexpected value: wanted %d ; got %d", input[0], v)
			}
		})
		t.Run("Fail", func(t *testing.T) {
			unset()
			if c.Idx(5) != 0 {
				t.Errorf("unexpected value: wanted %d ; got %d", 0, c.Idx(5))
			}
			set()
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
			if v != eof {
				t.Errorf("unexpected value: wanted %d ; got %d", input[10], v)
			}
		})
		t.Run("HitHead", func(t *testing.T) {
			v := c.Offset(-30)
			if v != eof {
				t.Errorf("unexpected value: wanted %d ; got %d", input[0], v)
			}
		})
		t.Run("Fail", func(t *testing.T) {
			unset()
			if c.Offset(5) != 0 {
				t.Errorf("unexpected value: wanted %d ; got %d", 0, c.Offset(5))
			}
			set()
		})
	})
	t.Run("Extract", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			v := c.Extract(5, 8)
			if len(v) != 3 {
				t.Errorf("unexpected slice length")
			}
			if v[0] != input[5] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[5], v[0])
			}
			if v[1] != input[6] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[6], v[1])
			}
			if v[2] != input[7] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[7], v[2])
			}
		})

		t.Run("HitTail", func(t *testing.T) {
			v := c.Extract(9, 15)
			if len(v) != 2 {
				t.Errorf("unexpected slice length")
			}
			if v[0] != input[9] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[9], v[0])
			}
			if v[1] != input[10] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[10], v[1])
			}
		})
		t.Run("HitHead", func(t *testing.T) {
			v := c.Extract(-3, 2)
			if len(v) != 2 {
				t.Errorf("unexpected slice length")
			}
			if v[0] != input[0] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[0], v[0])
			}
			if v[1] != input[1] {
				t.Errorf("unexpected value: wanted %d ; got %d", input[1], v[1])
			}
		})
		t.Run("Fail", func(t *testing.T) {
			unset()
			if len(c.Extract(5, 8)) != 0 {
				t.Errorf("unexpected value: wanted %d ; got %d", 0, len(c.Extract(5, 8)))
			}
			set()
		})
	})
}
