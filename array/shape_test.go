package array_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	array "github.com/jimmyskull/math/array"
)

func TestShape(t *testing.T) {
	nilShape := array.Shape(nil)
	assert.ErrorIs(t, array.ErrInvalidShapeSize, nilShape.Validate())

	zeroDim := array.Shape{}
	assert.Nil(t, zeroDim.Validate())

	singleDim := array.Shape{10}
	assert.Nil(t, singleDim.Validate())

	array2dDim := array.Shape{10, 20}
	assert.Nil(t, array2dDim.Validate())

	zeroSingleDim := array.Shape{-1}
	assert.ErrorIs(t, array.ErrInvalidShapeDim, zeroSingleDim.Validate())

	zero2Dim := array.Shape{2, -1}
	assert.ErrorIs(t, array.ErrInvalidShapeDim, zero2Dim.Validate())

	negNDim := array.Shape{10, 10, -2, 10}
	assert.ErrorIs(t, array.ErrInvalidShapeDim, negNDim.Validate())

	posNDim := array.Shape{10, 10, 1, 10}
	assert.Nil(t, posNDim.Validate())
}

// func TestShapeAndStrides(t *testing.T) {
// 	assert.Nil(t, array.Shape{0}.ValidateStrides(array.Strides{0}))
// 	assert.Nil(t, array.Shape{2}.ValidateStrides(array.Strides{0}))
// 	assert.Nil(t, array.Shape{2, 2}.ValidateStrides(array.Strides{0, 0}))
// 	assert.Nil(t, array.Shape{2, 0}.ValidateStrides(array.Strides{10, 100}))
// 	assert.Nil(t, array.Shape{2, 2}.ValidateStrides(array.Strides{2, 1}))
// 	assert.ErrorIs(
// 		t,
// 		array.Shape{2, 2}.ValidateStrides(array.Strides{3, 1}),
// 		array.ErrUnmatchedShapeAndStrides)
// 	assert.ErrorIs(
// 		t,
// 		array.Shape{10}.ValidateStrides(array.Strides{2}),
// 		array.ErrUnmatchedShapeAndStrides)
// }
