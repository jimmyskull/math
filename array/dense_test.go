package array_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jimmyskull/math/array"
)

func TestDenseNilShape(t *testing.T) {
	_, err := array.NewDense(nil, array.RowMajorLayout)
	assert.ErrorIs(t, err, array.ErrInvalidShapeSize)
	_, err = array.NewDenseWithStrides(nil, nil, array.DefaultAttributes)
	assert.ErrorIs(t, err, array.ErrInvalidShapeSize)
}

func TestDense0D(t *testing.T) {
	var arr *array.Dense
	var err error
	arr, err = array.NewDenseWithStrides(
		array.Shape{},
		array.Strides{},
		array.Contiguous|array.Writeable)
	if assert.Nil(t, err) {
		assert.Equal(t, array.Shape{}, arr.Shape)
		assert.Equal(t, array.Strides{}, arr.Strides)
		assert.Equal(t, 1, arr.Size())
		v, err := arr.Get(array.Indices{})
		assert.Nil(t, err)
		assert.Equal(t, 0.0, v)
		assert.Nil(t, arr.Set(array.Indices{}, 3))
		v, err = arr.Get(array.Indices{})
		assert.Nil(t, err)
		assert.Equal(t, 3.0, v)
		v, err = arr.Get(nil)
		assert.Nil(t, err)
		assert.Equal(t, 3.0, v)
		v, err = arr.Get(array.Indices{0})
		assert.Nil(t, err)
		assert.Equal(t, 3.0, v)
		v, err = arr.Get(array.Indices{1})
		if assert.Error(t, err) {
			serr := err.(*array.OutOfBoundsError)
			assert.Equal(t, 1, serr.Index)
			assert.Equal(t, -1, serr.Axis)
			assert.Equal(t, 1, serr.DimSize)
		}
		assert.True(t, math.IsNaN(v))

		v, err = arr.Get(array.Indices{0, 0})
		assert.ErrorIs(t, err, array.ErrIncorrectIndices)
		assert.True(t, math.IsNaN(v))
	}
	arr, err = array.NewDenseWithStrides(
		array.Shape{}, nil, array.Contiguous|array.Writeable)
	assert.ErrorIs(t, err, array.ErrInvalidStridesLength)
}

func TestDense1D(t *testing.T) {
	var arr *array.Dense
	var err error
	var v float64
	arr, err = array.NewDenseWithStrides(
		array.Shape{2},
		array.Strides{8},
		array.Contiguous|array.Writeable)
	if assert.Nil(t, err) {
		assert.Equal(t, array.Shape{2}, arr.Shape)
		assert.Equal(t, array.Strides{8}, arr.Strides)
		v, err = arr.Get(array.Indices{1})
		assert.Nil(t, err)
		assert.Equal(t, 0.0, v)
		assert.Nil(t, arr.Set(array.Indices{1}, 3))
		v, err = arr.Get(array.Indices{1})
		assert.Equal(t, 3.0, v)
		v, err = arr.Get(array.Indices{-1})
		assert.Nil(t, err)
		assert.Equal(t, 3.0, v)
		v, err = arr.Get(array.Indices{-2})
		assert.Nil(t, err)
		assert.Equal(t, 0.0, v)
	}
	arr, err = array.NewDenseWithStrides(
		array.Shape{2},
		array.Strides{8},
		array.Contiguous|array.Writeable)
	if assert.Nil(t, err) {
		assert.Equal(t, array.Shape{2}, arr.Shape)
		assert.Equal(t, array.Strides{8}, arr.Strides)
		v, err = arr.Get(array.Indices{})
		assert.ErrorIs(t, err, array.ErrIncorrectIndices)
		assert.True(t, math.IsNaN(v))
		v, err = arr.Get(array.Indices{0})
		assert.Nil(t, err)
		assert.Equal(t, 0.0, v)
		v, err = arr.Get(array.Indices{0, 0})
		assert.ErrorIs(t, err, array.ErrIncorrectIndices)
		assert.True(t, math.IsNaN(v))
	}
	arr, err = array.NewDenseWithStrides(
		array.Shape{2}, nil, array.Contiguous|array.Writeable)
	assert.ErrorIs(t, err, array.ErrInvalidStridesLength)
}

func TestDense2D(t *testing.T) {
	var arr *array.Dense
	var err error
	arr, err = array.NewDenseWithStrides(
		array.Shape{2, 2},
		array.Strides{16, 8},
		array.Contiguous|array.Writeable,
	)
	if assert.Nil(t, err) {
		assert.Equal(t, array.Shape{2, 2}, arr.Shape)
		assert.Equal(t, array.Strides{16, 8}, arr.Strides)
	}
	arr, err = array.NewDenseWithStrides(
		array.Shape{2, 2},
		array.Strides{16, 8},
		array.Contiguous|array.Writeable,
	)
	if assert.Nil(t, err) {
		assert.Equal(t, array.Shape{2, 2}, arr.Shape)
		assert.Equal(t, array.Strides{16, 8}, arr.Strides)
		v, err := arr.Get(array.Indices{1, 1})
		assert.Nil(t, err)
		assert.Equal(t, 0.0, v)
		assert.Nil(t, arr.Set(array.Indices{0, 0}, 1))
		assert.Nil(t, arr.Set(array.Indices{1, 1}, 3))
		v, err = arr.Get(array.Indices{1, 1})
		assert.Nil(t, err)
		assert.Equal(t, 3.0, v)
		v, err = arr.Get(array.Indices{1, -1})
		assert.Nil(t, err)
		assert.Equal(t, 3.0, v)
		v, err = arr.Get(array.Indices{-1, 1})
		assert.Nil(t, err)
		assert.Equal(t, 3.0, v)
		v, err = arr.Get(array.Indices{-1, -1})
		assert.Nil(t, err)
		assert.Equal(t, 3.0, v)
		v, err = arr.Get(array.Indices{-2, -2})
		assert.Nil(t, err)
		assert.Equal(t, 1.0, v)
		v, err = arr.Get(array.Indices{0, 0})
		assert.Nil(t, err)
		assert.Equal(t, 1.0, v)
		v, err = arr.Get(array.Indices{})
		assert.ErrorIs(t, err, array.ErrIncorrectIndices)
		assert.True(t, math.IsNaN(v))
		v, err = arr.Get(array.Indices{1})
		assert.ErrorIs(t, err, array.ErrIncorrectIndices)
		assert.True(t, math.IsNaN(v))
		v, err = arr.Get(array.Indices{0, 0, 0, 0})
		assert.ErrorIs(t, err, array.ErrIncorrectIndices)
		assert.True(t, math.IsNaN(v))
	}
	arr, err = array.NewDenseWithStrides(
		array.Shape{2, 2},
		nil,
		array.Contiguous|array.Writeable,
	)
	assert.ErrorIs(t, err, array.ErrInvalidStridesLength)
}
