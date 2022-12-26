# cur

*a generic Cursor interface for Go slices*

__________________

## Overview

`cur` is a generic Cursor interface for Go slices, allowing the caller to navigate through a slice in a controlled manner, freely moving forwards and backwards within a slice, and also to peek certain elements without actually moving the reference index.

## Features

The Cursor interface exposes a series of helper methods to navigate through a slice:

```go
// Cursor navigates through a slice in a controlled manner, allowing the
// caller to move forward, backwards, and jump around the slice as they need
type Cursor[T any] interface {

	// Cur returns the same indexed item in the slice
	Cur() T

	// Pos returns the current position in the cursor
	Pos() int

	// Len returns the total size of the underlying slice
	Len() int

	// Next returns the next item in the slice, or the zero-value for T as EOF
	Next() T

	// Prev returns the previous item in the slice, or the zero-value for T as EOF if
	// index is / would be less than zero
	Prev() T

	// Peek returns the next indexed item without advancing the cursor
	//
	// If the next token overflows the slice, returns the zero-value for T as EOF
	Peek() T

	// Head returns to the beginning of the slice
	Head() T

	// Tail jumps to the end of the slice
	Tail() T

	// Idx jumps to the specific index `idx` in the slice
	//
	// If the input index is below 0, the zero-value for T as EOF
	// If the input index is greater than the size of the slice, the zero-value for T as EOF
	Idx(idx int) T

	// Offset advances or rewinds `amount` steps in the slice, be it a positive or negative
	// input.
	//
	// If the result offset is below 0, the zero-value for T as EOF
	// If the result offset is greater than the size of the slice, the zero-value for T as EOF
	Offset(amount int) T

	// PeekIdx returns the next indexed item without advancing the cursor,
	// with the index `idx`
	//
	// If the input index is below 0, the zero-value for T as EOF
	// If the input index is greater than the size of the slice, the zero-value for T as EOF
	PeekIdx(idx int) T

	// PeekOffset returns the next indexed item without advancing the cursor,
	// with offset `amount`
	//
	// If the result offset is below 0, the zero-value for T as EOF
	// If the result offset is greater than the size of the slice, the zero-value for T as EOF
	PeekOffset(amount int) T

	// Extract returns a slice from index `start` to index `end`
	Extract(start, end int) []T
}
```

Creating a cursor only takes an input slice, provided that it has one or more elements in it. An empty slice will return a `nil` Cursor:

```go
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
```