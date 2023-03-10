package cur

type ptrCursor[T any] struct {
	slice *[]T
	pos   int
}

// NewCursor returns a Cursor for the input slice, or nil if the slice is empty
func Ptr[T any](slice *[]T) Cursor[T] {
	if slice == nil {
		return nil
	}
	return &ptrCursor[T]{
		slice: slice,
	}
}

// Cur returns the same indexed item in the slice
func (c *ptrCursor[T]) Cur() T {
	if c.slice == nil || c.pos >= len(*c.slice) {
		var zero T
		return zero
	}
	s := *c.slice
	return s[c.pos]
}

// Pos returns the current position in the cursor
func (c *ptrCursor[T]) Pos() int {
	if c.slice == nil || c.pos >= len(*c.slice) {
		return -1
	}
	return c.pos
}

// Len returns the total size of the underlying slice
func (c *ptrCursor[T]) Len() int {
	if c.slice == nil {
		return -1
	}
	return len(*c.slice)
}

// Next returns the next item in the slice, or the zero-value for T as EOF
func (c *ptrCursor[T]) Next() T {
	if c.slice == nil || c.pos >= len(*c.slice) {
		var zero T
		return zero
	}
	c.pos++
	s := *c.slice
	return s[c.pos-1]
}

// Prev returns the previous item in the slice, or the zero-value for T as EOF if
// index is / would be less than zero
func (c *ptrCursor[T]) Prev() T {
	if c.slice == nil || c.pos <= 0 {
		var zero T
		return zero
	}
	c.pos--
	s := *c.slice
	return s[c.pos]
}

// Peek returns the next indexed item without advancing the cursor
//
// If the next token overflows the slice, returns the zero-value for T as EOF
func (c *ptrCursor[T]) Peek() T {
	if c.slice == nil || c.pos+1 >= len(*c.slice) {
		var eof T
		return eof
	}
	s := *c.slice
	return s[c.pos+1]
}

// Head returns to the beginning of the slice
func (c *ptrCursor[T]) Head() T {
	if c.slice == nil || len(*c.slice) == 0 {
		var zero T
		return zero
	}
	c.pos = 0
	return c.Next()
}

// Tail jumps to the end of the slice
func (c *ptrCursor[T]) Tail() T {
	if c.slice == nil || len(*c.slice) == 0 {
		var zero T
		return zero
	}
	c.pos = len(*c.slice)
	return c.Prev()
}

// Idx jumps to the specific index `idx` in the slice
//
// If the input index is below 0, the zero-value for T as EOF
// If the input index is greater than the size of the slice, the zero-value for T as EOF
func (c *ptrCursor[T]) Idx(idx int) T {
	if c.slice == nil || idx < 0 || idx >= len(*c.slice) {
		var zero T
		return zero
	}

	c.pos = idx
	s := *c.slice
	return s[idx]
}

// Offset advances or rewinds `amount` steps in the slice, be it a positive or negative
// input.
//
// If the result offset is below 0, the zero-value for T as EOF
// If the result offset is greater than the size of the slice, the zero-value for T as EOF
func (c *ptrCursor[T]) Offset(amount int) T {
	if c.slice == nil || c.pos+amount < 0 || c.pos+amount >= len(*c.slice) {
		var zero T
		return zero
	}

	c.pos += amount
	s := *c.slice
	return s[c.pos]
}

// PeekIdx returns the next indexed item without advancing the cursor,
// with the index `idx`
//
// If the input index is below 0, the zero-value for T as EOF
// If the input index is greater than the size of the slice, the zero-value for T as EOF
func (c *ptrCursor[T]) PeekIdx(idx int) T {
	if c.slice == nil || idx < 0 || idx >= len(*c.slice)-1 {
		var zero T
		return zero
	}

	s := *c.slice
	return s[idx]
}

// PeekOffset returns the next indexed item without advancing the cursor,
// with offset `amount`
//
// If the result offset is below 0, the zero-value for T as EOF
// If the result offset is greater than the size of the slice, the zero-value for T as EOF
func (c *ptrCursor[T]) PeekOffset(amount int) T {
	if c.slice == nil || c.pos+amount < 0 || c.pos+amount >= len(*c.slice)-1 {
		var zero T
		return zero
	}

	s := *c.slice
	return s[c.pos+amount]
}

// Extract returns a slice from index `start` to index `end`
func (c *ptrCursor[T]) Extract(start, end int) []T {
	if c.slice == nil {
		return nil
	}
	if start < 0 {
		start = 0
	}
	if end > len(*c.slice) {
		end = len(*c.slice)
	}
	for start > end {
		start--
	}

	s := *c.slice
	return s[start:end]
}
