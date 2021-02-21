package array

import (
	"fmt"
)

// Strides is a tuple of bytes to step in each dimension when
// traversing an array.
type Strides []int

// NewStrides returns an initialized strides according to an
// array's shape and the specified memory layout.
func NewStrides(shape Shape, dtype DType, attrs Attributes) (Strides, error) {
	if err := shape.Validate(); err != nil {
		return nil, err
	}
	switch {
	case attrs.Is(RowMajorLayout | ColumnMajorLayout):
		return nil, ErrInvalidLayout
	case attrs.Is(RowMajorLayout):
		return newStridesWithRowMajorLayout(shape, dtype), nil
	case attrs.Is(ColumnMajorLayout):
		return newStridesWithColumnMajorLayout(shape, dtype), nil
	default:
		return nil, ErrInvalidLayout
	}
}

// CanIndex returns an error when the strides would map an array with
// a given shape and data-type outside a contiguous slice with the
// same size of the number of elements.
func (s Strides) CanIndex(shape Shape, dtype DType) error {
	var elements, nbytes int
	var err error
	elements, err = shape.Size()
	elementsSize := elements * dtype.Size()
	if err != nil {
		return err
	}
	nbytes = s.Size(shape)
	if elementsSize < nbytes {
		return ErrUnmatchedShapeAndStrides
	}
	return nil
}

// Size returns the total number of bytes mapped by the strides with
// a given shape and data-type.
func (s Strides) Size(shape Shape) int {
	size := 0
	for i, stride := range s {
		if shape[i] == 0 {
			return 0
		}
		size += stride * (shape[i] - 1)
	}
	return (size + 1)
}

func (s Strides) String() string {
	return fmt.Sprintf("Strides(%s)", sprintIntSliceWithSep(", ", s))
}

func newStridesWithRowMajorLayout(shape Shape, dtype DType) Strides {
	n := len(shape)
	s := make(Strides, n)
	offset := dtype.Size()
	for i := n - 1; i >= 0; i-- {
		s[i] = offset
		offset *= shape[i]
	}
	return s
}

func newStridesWithColumnMajorLayout(shape Shape, dtype DType) Strides {
	n := len(shape)
	s := make(Strides, n)
	offset := dtype.Size()
	for i := 0; i < n; i++ {
		s[i] = offset
		offset *= shape[i]
	}
	return s
}
