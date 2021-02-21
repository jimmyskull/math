package array_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jimmyskull/math/array"
)

func TestStride1D(t *testing.T) {
	shape := array.Shape{10}
	s, err := array.NewStrides(shape, array.Float64, array.RowMajorLayout)
	assert.Nil(t, err)
	assert.Equal(t, array.Strides{8}, s)
	s, err = array.NewStrides(shape, array.Float64, array.ColumnMajorLayout)
	assert.Nil(t, err)
	assert.Equal(t, array.Strides{8}, s)
	s, err = array.NewStrides(
		shape,
		array.Float64,
		array.ColumnMajorLayout|array.RowMajorLayout,
	)
	assert.ErrorIs(t, err, array.ErrInvalidLayout)
}

func TestStrideND(t *testing.T) {
	shape := array.Shape{5, 4, 3, 2, 1}
	s, err := array.NewStrides(shape, array.Float64, array.RowMajorLayout)
	assert.Nil(t, err)
	assert.Equal(t, array.Strides{192, 48, 16, 8, 8}, s)
	s, err = array.NewStrides(shape, array.Float64, array.ColumnMajorLayout)
	assert.Nil(t, err)
	assert.Equal(t, array.Strides{8, 40, 160, 480, 960}, s)
}

func TestStridesCanIndex(t *testing.T) {
	var strides array.Strides
	var shape array.Shape
	err := strides.CanIndex(shape, array.Float64)
	assert.Nil(t, err)
}
