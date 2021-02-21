package array

import (
	"math"

	"github.com/jimmyskull/math/scalar"
)

// Arange returns a Dense array with values evenly spaced within a given
// interval.
//
// Values are generated within the half-open interval `[start, stop)`.
func Arange(start, stop, step float64) (*Dense, error) {
	length, err := arangeLength(start, stop, step)
	if err != nil {
		return nil, err
	}
	d, err := NewDense(Shape{length}, DefaultAttributes)
	if err != nil {
		return nil, err
	}
	d.Fill(start, step)
	return d, nil
}

func arangeLength(start, stop, step float64) (int, error) {
	delta := stop - start
	length := delta / step
	// Underflow check.
	if length == 0.0 && delta != 0.0 {
		if math.Signbit(length) {
			length = 0
		} else {
			length = 1
		}
	} else {
		// Division-by-zero check.
		if step == 0.0 {
			return 0, ErrZeroDivision
		}
	}
	length = math.Ceil(length)
	if math.IsInf(length, 0) || math.IsNaN(length) {
		return 0, &Error{
			Operation: "arange",
			Message:   "cannot compute length",
		}
	}
	// Check if casting from float64 to int would be out-of-bounds.
	if length < float64(scalar.MinInt) || length > float64(scalar.MaxInt) {
		return 0, &Error{
			Operation: "arange",
			Message:   "overflow while computing length",
		}
	}
	if length <= 0 {
		length = 0
	}
	return int(length), nil
}
