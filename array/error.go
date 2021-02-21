package array

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidShapeSize is returned when an array shape is either too
	// small or too large.
	ErrInvalidShapeSize = fmt.Errorf(
		"number of dimensions must be within [%d, %d]",
		MinDimensions, MaxDimensions)

	// ErrInvalidShapeDim is returned when an array contains a dimension
	// with an invalid size.
	ErrInvalidShapeDim = errors.New("negative dimensions are not allowed")

	// ErrUnmatchedShapeAndStrides is returned when the provided strides
	// are incompatible, such as in the case of generating an
	// non-contiguous array.
	ErrUnmatchedShapeAndStrides = errors.New(
		"strides is incompatible with shape")

	// ErrInvalidLayout happens when un unsupported layout scheme is
	// specified.
	ErrInvalidLayout = errors.New("invalid array layout")

	// ErrInvalidStridesLength is returned when the given strides are nil
	// or are incompatible with the shape specification.
	ErrInvalidStridesLength = errors.New(
		"strides must be the same length as shape")

	// ErrIncorrectIndices is returned when the number of indices to access
	// an item in an array is insufficient to determine which item should
	// be accessed.
	ErrIncorrectIndices = errors.New("incorrect number of indices for array")

	// ErrZeroDivision is returned when a division by zero occurred in
	// some specific situations, such as in Arange.
	ErrZeroDivision = errors.New("division by zero")
)

// Error is used to describe errors that are not otherwise
// specified.
type Error struct {
	Operation string
	Message   string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Operation, e.Message)
}

// OutOfBoundsError describes an invalid acces for an array item.
type OutOfBoundsError struct {
	Index   int
	Axis    int
	DimSize int
}

func (e *OutOfBoundsError) Error() string {
	if e.Axis < 0 {
		msg := fmt.Sprintf(
			"index %d is out of bounds for size %d", e.Index, e.DimSize)
		return msg
	}
	msg := fmt.Sprintf(
		"index %d is out of bounds for axis %d with size %d",
		e.Index, e.Axis, e.DimSize)
	return msg
}
