package array

import "strings"

// Attributes specifies the state of the array.
type Attributes int

const (
	// Contiguous indicates whether the array is contiguously disposed.
	Contiguous Attributes = 1 << iota

	// Writeable indicates whether the array is mutable or not.
	Writeable

	// RowMajorLayout defines an array with elements disposed row-by-row
	// in a contiguous single-segment, memory layout.
	RowMajorLayout

	// ColumnMajorLayout defines an array with elements disposed
	// column-by-column in a contiguous single-segment, memory layout.
	ColumnMajorLayout
)

// DefaultAttributes can be used to set normally used attributes for a
// new array.
const DefaultAttributes = Contiguous | Writeable | ColumnMajorLayout

// Is returns whether the attribute set is a super set of the requested
// attributes.
func (a Attributes) Is(test Attributes) bool {
	if a&test == test {
		return true
	}
	return false
}

func (a Attributes) String() string {
	var attrs []string
	if a.Is(Contiguous) {
		attrs = append(attrs, "Contiguous")
	}
	if a.Is(Writeable) {
		attrs = append(attrs, "Writeable")
	}
	if a.Is(RowMajorLayout) {
		attrs = append(attrs, "RowMajorLayout")
	}
	if a.Is(ColumnMajorLayout) {
		attrs = append(attrs, "ColumnMajorLayout")
	}
	return strings.Join(attrs, ", ")
}
