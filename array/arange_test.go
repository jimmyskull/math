package array_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jimmyskull/math/array"
	"github.com/jimmyskull/math/scalar"
)

func TestArange(t *testing.T) {
	var arr *array.Dense
	var err error
	var v float64
	// Zero step returns a division-by-zero error.
	arr, err = array.Arange(0, 100, 0)
	if assert.ErrorIs(t, err, array.ErrZeroDivision) {
		assert.Nil(t, arr)
	}
	// Negative step creates an empty 1d array.
	arr, err = array.Arange(0, 100, -1)
	if assert.Nil(t, err) && assert.NotNil(t, arr) {
		assert.Equal(t, array.Shape{0}, arr.Shape)
		assert.Equal(t, array.Strides{8}, arr.Strides)
		assert.Equal(t, 0, arr.Size())
	}
	// Large positive step creates a single-value 1d array.
	arr, err = array.Arange(5, 10, 5)
	if assert.Nil(t, err) && assert.NotNil(t, arr) {
		assert.Equal(t, array.Shape{1}, arr.Shape)
		assert.Equal(t, array.Strides{8}, arr.Strides)
		assert.Equal(t, 1, arr.Size())
		v, err = arr.Get(array.Indices{0})
		assert.Nil(t, err)
		assert.Equal(t, 5.0, v)
	}
	arr, err = array.Arange(0, 1, 0.1)
	if assert.Nil(t, err) && assert.NotNil(t, arr) {
		assert.Equal(t, array.Shape{10}, arr.Shape)
		assert.Equal(t, array.Strides{8}, arr.Strides)
		assert.Equal(t, 10, arr.Size())
		expected := []float64{0., .1, .2, .3, .4, .5, .6, .7, .8, .9}
		for i := 0; i < 10; i++ {
			assert.InDelta(t, expected[i], arr.Data[i], 1e-10)
		}
	}
	arr, err = array.Arange(1, 0, -0.1)
	if assert.Nil(t, err) && assert.NotNil(t, arr) {
		assert.Equal(t, array.Shape{10}, arr.Shape)
		assert.Equal(t, array.Strides{8}, arr.Strides)
		assert.Equal(t, 10, arr.Size())
		expected := []float64{1., .9, .8, .7, .6, .5, .4, .3, .2, .1}
		for i := 0; i < 10; i++ {
			assert.InDelta(t, expected[i], arr.Data[i], 1e-10)
		}
	}
	arr, err = array.Arange(-1, 0, -0.1)
	if assert.Nil(t, err) && assert.NotNil(t, arr) {
		assert.Equal(t, array.Shape{0}, arr.Shape)
		assert.Equal(t, array.Strides{8}, arr.Strides)
		assert.Equal(t, 0, arr.Size())
	}
	arr, err = array.Arange(0, float64(scalar.MaxInt), 0.0001)
	if assert.Error(t, err) {
		serr := err.(*array.Error)
		assert.Equal(t, "arange", serr.Operation)
		assert.Equal(t, "overflow while computing length", serr.Message)
	}
	arr, err = array.Arange(1, 1, 0.1)
	if assert.Nil(t, err) && assert.NotNil(t, arr) {
		assert.Equal(t, array.Shape{0}, arr.Shape)
		assert.Equal(t, array.Strides{8}, arr.Strides)
		assert.Equal(t, 0, arr.Size())
	}
	// Test underflow handling.
	arr, err = array.Arange(1e-200, 1.1e-200, 1e200)
	if assert.Nil(t, err) && assert.NotNil(t, arr) {
		assert.Equal(t, array.Shape{1}, arr.Shape)
		assert.Equal(t, array.Strides{8}, arr.Strides)
		assert.Equal(t, 1, arr.Size())
	}
	arr, err = array.Arange(1.1e-200, 1e-200, 1e200)
	if assert.Nil(t, err) && assert.NotNil(t, arr) {
		assert.Equal(t, array.Shape{0}, arr.Shape)
		assert.Equal(t, array.Strides{8}, arr.Strides)
		assert.Equal(t, 0, arr.Size())
	}
	// Test overflow on length computation.
	arr, err = array.Arange(-1e200, 1e200, 1e-200)
	if assert.Error(t, err) {
		serr := err.(*array.Error)
		assert.Equal(t, "arange", serr.Operation)
		assert.Equal(t, "cannot compute length", serr.Message)
	}
}
