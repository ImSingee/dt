package dt

import (
	"fmt"
	"strconv"
)

type Stringer = fmt.Stringer

type String string

func (s String) String() string {
	return string(s)
}

func ToString(v interface{}) string {
	switch vv := v.(type) {
	case string:
		return vv
	case Stringer: // 包括 *GenericNumber
		return vv.String()
	case byte:
		return string(vv)
	case rune:
		return string(vv)
	case float64:
		return strconv.FormatFloat(vv, 'f', -1, 64)
	default:
		return fmt.Sprintf("%#+v", v)
	}
}

func ToStringer(v interface{}) Stringer {
	switch vv := v.(type) {
	case Stringer:
		return vv
	default:
		return String(ToString(v))
	}
}
