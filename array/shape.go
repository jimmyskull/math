package array

import (
	"fmt"

	"github.com/jimmyskull/math/scalar"
)

// Shape is a tuple of array dimensions.
type Shape []int

// MinDimensions limits the minimum number of components allowed in any
// multi-dimensional array.
const MinDimensions = 0

// MaxDimensions limits the number of component dimensions allowed in
// any multi-dimensional array.  Note, however, that a specific array
// implementation, such as of a space matrix, may have a lower
// dimensional limit.
const MaxDimensions = 32

// Size returns the number of elements mapped by this shape.  An error
// is returned if an overflow occurs when multiplying the dimensions
// together.
func (s Shape) Size() (int, error) {
	var err error
	size := 1
	for _, dim := range s {
		size, err = scalar.SafeMultiplyInt(size, dim)
		if err != nil {
			return size, err
		}
	}
	return size, nil
}

func (s Shape) String() string {
	return fmt.Sprintf("Shape(%s)", sprintIntSliceWithSep(", ", s))
}

// Validate returns an error when the shape does not meet requirements.
func (s Shape) Validate() error {
	if !s.validDimensions() {
		return ErrInvalidShapeSize
	}
	if !s.allNonNegative() {
		return ErrInvalidShapeDim
	}
	return nil
}

func (s Shape) allNonNegative() bool {
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] < 0 {
			return false
		}
	}
	return true
}

func (s Shape) validDimensions() bool {
	if s == nil {
		return false
	}
	nd := len(s)
	return nd >= MinDimensions && nd <= MaxDimensions
}
