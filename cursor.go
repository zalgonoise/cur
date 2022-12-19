package codegraph

// Cursor navigates through a slice in a controlled manner, allowing the
// caller to move forward, backwards, and jump around the slice as they need
type Cursor[T any] interface {
	// Cur returns the same indexed item in the slice
	Cur() T
	// Pos returns the current position in the cursor
	Pos() int
	// Len returns the total size of the underlying slice
	Len() int
	// Next returns the next item in the slice, or the last index if already
	// at the tail
	Next() T
	// Prev returns the previous item in the slice, or the first index if
	// already at the head
	Prev() T
	// Peek returns the next indexed item without advancing the cursor
	Peek() T
	// Head returns to the beginning of the slice
	Head() T
	// Tail jumps to the end of the slice
	Tail() T
	// Idx jumps to the specific index `idx` in the slice
	//
	// If the input index is below 0, the head of the slice is returned;
	// If the input index is greater than the size of the slice, the tail of the slice
	// is returned
	Idx(idx int) T
	// Offset advances or rewinds `amount` steps in the slice, be it a positive or negative
	// input.
	//
	// If the result offset is below 0, the head of the slice is returned;
	// If the result offset is greater than the size of the slice, the tail of the slice
	// is returned
	Offset(amount int) T
	// PeekIdx returns the next indexed item without advancing the cursor,
	// with the index `idx`
	PeekIdx(idx int) T
	// PeekOffset returns the next indexed item without advancing the cursor,
	// with offset `amount`
	PeekOffset(amount int) T
}

type cursor[T any] struct {
	slice []T
	idx   int
}

// NewCursor returns a Cursor for the input slice, or nil if the slice is empty
func NewCursor[T any](slice []T) Cursor[T] {
	if len(slice) == 0 {
		return nil
	}
	return &cursor[T]{
		slice: slice,
	}
}

// Cur returns the same indexed item in the slice
func (c *cursor[T]) Cur() T {
	return c.slice[c.idx]
}

// Pos returns the current position in the cursor
func (c *cursor[T]) Pos() int {
	return c.idx
}

// Len returns the total size of the underlying slice
func (c *cursor[T]) Len() int {
	return len(c.slice)
}

// Next returns the next item in the slice, or the last index if already
// at the tail
func (c *cursor[T]) Next() T {
	if c.idx < len(c.slice)-1 {
		c.idx++
	}
	return c.slice[c.idx]
}

// Prev returns the previous item in the slice, or the first index if
// already at the head
func (c *cursor[T]) Prev() T {
	if c.idx == 0 {
		return c.slice[0]
	}
	c.idx--
	return c.slice[c.idx]
}

// Peek returns the next indexed item without advancing the cursor
func (c *cursor[T]) Peek() T {
	if c.idx+1 > len(c.slice)-1 {
		return c.slice[c.idx]
	}
	return c.slice[c.idx+1]
}

// Head returns to the beginning of the slice
func (c *cursor[T]) Head() T {
	c.idx = 0
	return c.slice[c.idx]
}

// Tail jumps to the end of the slice
func (c *cursor[T]) Tail() T {
	c.idx = len(c.slice) - 1
	return c.slice[c.idx]
}

// Idx jumps to the specific index `idx` in the slice
//
// If the input index is below 0, the head of the slice is returned;
// If the input index is greater than the size of the slice, the tail of the slice
// is returned
func (c *cursor[T]) Idx(idx int) T {
	if idx < 0 {
		return c.Head()
	}
	if idx >= len(c.slice) {
		return c.Tail()
	}
	c.idx = idx
	return c.slice[c.idx]
}

// Offset advances or rewinds `amount` steps in the slice, be it a positive or negative
// input.
//
// If the result offset is below 0, the head of the slice is returned;
// If the result offset is greater than the size of the slice, the tail of the slice
// is returned
func (c *cursor[T]) Offset(amount int) T {
	if c.idx+amount < 0 {
		return c.Head()
	}
	if c.idx+amount >= len(c.slice) {
		return c.Tail()
	}
	c.idx += amount
	return c.slice[c.idx]
}

// PeekIdx returns the next indexed item without advancing the cursor,
// with the index `idx`
func (c *cursor[T]) PeekIdx(idx int) T {
	if idx < 0 {
		return c.slice[0]
	}
	if idx >= len(c.slice)-1 {
		return c.slice[len(c.slice)-1]
	}
	return c.slice[idx]
}

// PeekOffset returns the next indexed item without advancing the cursor,
// with offset `amount`
func (c *cursor[T]) PeekOffset(amount int) T {
	if c.idx+amount < 0 {
		return c.slice[0]
	}
	if c.idx+amount >= len(c.slice)-1 {
		return c.slice[len(c.slice)-1]
	}
	return c.slice[c.idx+amount]
}
