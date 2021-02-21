package scalar_test

import (
	"math"
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/jimmyskull/math/scalar"
	"github.com/stretchr/testify/assert"
)

func TestMultiplyInt64(t *testing.T) {
	f := fuzz.New()
	var a, r int64
	var err error
	for i := 0; i < 50000; i++ {
		f.Fuzz(&a)
		r, err = scalar.SafeMultiplyInt64(a, 1)
		assert.Equal(t, a, r)
		assert.Nil(t, err)
		r, err = scalar.SafeMultiplyInt64(1, a)
		assert.Equal(t, a, r)
		assert.Nil(t, err)
	}
	overflowErr := &scalar.Overflow{}
	maxForDoubling := int64(math.MaxInt64 / 2)
	minForDoubling := int64(math.MinInt64 / 2)
	for i := 0; i < 50000; i++ {
		f.Fuzz(&a)
		_, err = scalar.SafeMultiplyInt64(a, math.MaxInt64)
		assert.IsType(t, overflowErr, err)
		_, err = scalar.SafeMultiplyInt64(math.MaxInt64, a)
		assert.IsType(t, overflowErr, err)
		if a >= maxForDoubling || a <= minForDoubling {
			a = a / 2
		}
		r, err = scalar.SafeMultiplyInt64(a, 2)
		assert.Equal(t, a+a, r)
		assert.Nil(t, err)
	}
}
