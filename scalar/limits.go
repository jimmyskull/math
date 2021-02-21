package scalar

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

// SafeMultiplyInt returns the multiplication of two int, indicating
// an error such as an overflow.
func SafeMultiplyInt(a, b int) (int, error) {
	result := a * b
	if a == 0 || b == 0 || a == 1 || b == 1 {
		return result, nil
	}
	if a == MinInt || b == MinInt {
		return result, NewMultiplyOverflowError(a, b)
	}
	if result/b != a {
		return result, NewMultiplyOverflowError(a, b)
	}
	return result, nil
}
