// Package array provides an n-dimensional array type, the Dense, which
// describes an indexed collection of items of the same data type.
package array

import (
	"math"
)

// Dense is a homogeneous multi-dimensional array of float64 items.
//
// Dense arrays may have from 0 to 32 dimensions.
//
// A zero-dimensional array is a scalar item with Shape{} and Strides{}.
//
// An one-dimensional array is a vector with Shape{n} for n items, with
// Strides{8} for the 8-byte float64 data type.
//
// A two-dimensional array is a matrix with Shape{m, n} for m rows and
// n columns, with either Strides{8*n, 8} for row-major layout or
// Strides{8, 8*m} for column-major layout.
type Dense struct {
	Data    []float64
	DType   DType
	Shape   Shape
	Strides Strides
	Attrs   Attributes
}

// NewDense returns a new float64, multi-dimensional array with a given
// shape, with elements disposed with the specified strides.
func NewDense(shape Shape, attrs Attributes) (*Dense, error) {
	newStrides, err := NewStrides(shape, Float64, attrs)
	if err != nil {
		return nil, err
	}
	return NewDenseWithStrides(shape, newStrides, attrs)
}

// NewDenseWithStrides returns a new float64, multi-dimensional array
// with a given shape, with elements disposed with the specified
// strides.  Although possible to manually define strides, this is not
// recommended.
func NewDenseWithStrides(
	shape Shape, strides Strides, attrs Attributes,
) (*Dense, error) {
	var err error
	err = shape.Validate()
	if err != nil {
		return nil, err
	}
	if strides == nil {
		return nil, ErrInvalidStridesLength
	}
	var elements int
	err = strides.CanIndex(shape, Float64)
	if err != nil {
		return nil, err
	}
	elements, err = shape.Size()
	data := make([]float64, elements)
	return &Dense{
		Data:    data,
		DType:   Float64,
		Shape:   shape,
		Strides: strides,
		Attrs:   attrs,
	}, err
}

// Fill sets all items in array with start+delta*i for each item i.
func (d *Dense) Fill(start, delta float64) {
	n := d.Size()
	if delta == 0.0 {
		// Assign a constant scalar for all values.
		for i := 0; i < n; i++ {
			d.Data[i] = start
		}
	} else {
		// Assign stepped scalar for all values.
		for i := 0; i < n; i++ {
			d.Data[i] = start + float64(i)*delta
		}
	}
}

// Get returns the item at a given position.
func (d *Dense) Get(indices Indices) (float64, error) {
	offset, err := d.Offset(indices)
	if err != nil {
		return math.NaN(), err
	}
	return d.Data[offset], nil
}

// Set replaces an item at a given position.
func (d *Dense) Set(indices Indices, item float64) error {
	offset, err := d.Offset(indices)
	if err != nil {
		return err
	}
	d.Data[offset] = item
	return nil
}

// Size returns the number of items in the array.
func (d *Dense) Size() int {
	// We ignore errors here because an array cannot be created without
	// validating the shape's size.
	s, _ := d.Shape.Size()
	return s
}

// Offset returns the position of an index in the data.
func (d *Dense) Offset(indices Indices) (int, error) {
	pos := 0
	if len(d.Shape) == 0 {
		switch len(indices) {
		case 0:
			return pos, nil
		case 1:
			idx, err := adjustedIndex(indices[0], -1, 1)
			if err != nil {
				return 0, err
			}
			pos = idx * d.DType.Size()
		default:
			return 0, ErrIncorrectIndices
		}
	} else {
		if len(d.Shape) != len(indices) {
			return 0, ErrIncorrectIndices
		}
	}
	for axis, dimSize := range d.Shape {
		adjIndex, err := adjustedIndex(indices[axis], axis, dimSize)
		if err != nil {
			return 0, err
		}
		pos += adjIndex * d.Strides[axis]
	}
	return pos / d.DType.Size(), nil
}

func adjustedIndex(index int, axis int, dimSize int) (int, error) {
	if index < -dimSize || index >= dimSize {
		return 0, &OutOfBoundsError{Index: index, Axis: axis, DimSize: dimSize}
	}
	if index < 0 {
		return index + dimSize, nil
	}
	return index, nil
}
