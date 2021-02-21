package scalar

import "fmt"

// Overflow indicates an overflow ocurred when operating two numbers.
type Overflow struct {
	lhs interface{}
	rhs interface{}
	msg string
}

func (e *Overflow) Error() string {
	return e.msg
}

// NewMultiplyOverflowError returns an error specification of
// multiplication overflow.
func NewMultiplyOverflowError(lhs interface{}, rhs interface{}) *Overflow {
	return &Overflow{
		lhs: lhs,
		rhs: rhs,
		msg: fmt.Sprintf("overflow multiplying %v and %v", lhs, rhs),
	}
}
