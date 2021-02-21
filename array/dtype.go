package array

import (
	"fmt"
	"unsafe"
)

// DType is the data type for an array.
type DType string

const (
	// Float64 is the Go's float64 type.
	Float64 DType = "float64"
)

// Size returns the number of bytes used to store an element of this
// data-type.
func (d DType) Size() int {
	switch d {
	case Float64:
		return int(unsafe.Sizeof(float64(0)))
	}
	panic(fmt.Sprintf("array: invalid dtype %#v", d))
}
