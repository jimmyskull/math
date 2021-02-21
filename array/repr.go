package array

import (
	"fmt"
	"strings"
)

func sprintIntSliceWithSep(sep string, values []int) string {
	var repr strings.Builder
	for i, value := range values {
		if i > 0 {
			repr.WriteString(sep)
		}
		repr.WriteString(fmt.Sprint(value))
	}
	return repr.String()
}
