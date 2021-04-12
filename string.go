package dt

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

func ToString(v interface{}) string {
	switch vv := v.(type) {
	case string:
		return vv
	case Stringer: // 包括 *GenericNumber
		return vv.String()
	case float64:
		return strconv.FormatFloat(vv, 'f', -1, 64)
	default:
		return fmt.Sprintf("%#+v", v)
	}
}
