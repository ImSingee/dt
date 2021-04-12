package dt

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

var IsNotNumber = fmt.Errorf("is not a number")

type Number interface {
	AsNumber() *GenericNumber
}

type GenericNumber struct {
	literal string

	above64bit bool // 是否不足以使用 64 位表示 (number 类型为 *big.Float)
	float      bool // 是否为浮点数（number 类型为 float64）
	unsigned   bool // 是否有符号（number 类型为 uint64）

	number interface{} // 类型为 int64/uint64/float64/*big.Float
}

func convertSignedIntToInt64(num interface{}) int64 {
	switch v := num.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	default:
		return 0
	}
}

func convertUnsignedIntToUInt64(num interface{}) uint64 {
	switch v := num.(type) {
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return v
	default:
		return 0
	}
}

func convertFloatToFloat64(num interface{}) float64 {
	switch v := num.(type) {
	case float32:
		return float64(v)
	case float64:
		return v
	default:
		return 0
	}
}

type stringer interface {
	String() string
}

func convertStringerToString(num interface{}) string {
	if v, ok := num.(string); ok {
		return v
	}
	if v, ok := num.(stringer); ok {
		return v.String()
	}

	return ""
}

// 从字符串获取对应表示的数字值
func numberFromString(num string) (*GenericNumber, error) {
	// 尝试整数
	v, err := strconv.ParseInt(num, 10, 64)
	if err == nil {
		return &GenericNumber{
			literal:    num,
			float:      false,
			unsigned:   false,
			above64bit: false,
			number:     v,
		}, nil
	}
	e := err.(*strconv.NumError)
	if errors.Is(e.Err, strconv.ErrSyntax) {
		// 不是整数，尝试浮点数
		v, err := strconv.ParseFloat(num, 64)
		if err == nil {
			return &GenericNumber{
				literal:    num,
				float:      true,
				unsigned:   false,
				above64bit: false,
				number:     v,
			}, nil
		}

		// 返回错误 （不是数字）
		return nil, IsNotNumber
	}

	// 尝试正整数
	vv, err := strconv.ParseUint(num, 10, 64)
	if err == nil {
		return &GenericNumber{
			literal:    num,
			float:      false,
			unsigned:   true,
			above64bit: false,
			number:     vv,
		}, nil
	}

	// 使用大整数
	f, _, err := new(big.Float).SetPrec(big.MaxPrec).Parse(num, 10)
	if err != nil {
		return nil, IsNotNumber
	}

	return &GenericNumber{
		literal:    num,
		above64bit: true,
		number:     f,
	}, nil
}

func newGenericNumber(num interface{}) (*GenericNumber, error) {
	switch v := num.(type) {
	case *GenericNumber:
		return v, nil
	case int, int8, int16, int32, int64:
		vv := convertSignedIntToInt64(v)
		return &GenericNumber{
			literal:    strconv.FormatInt(vv, 10),
			float:      false,
			unsigned:   false,
			above64bit: false,
			number:     vv,
		}, nil
	case uint, uint8, uint16, uint32, uint64:
		vv := convertUnsignedIntToUInt64(v)
		return &GenericNumber{
			literal:    strconv.FormatUint(vv, 10),
			float:      false,
			unsigned:   true,
			above64bit: false,
			number:     vv,
		}, nil
	case float32, float64:
		vv := convertFloatToFloat64(v)
		return &GenericNumber{
			literal:    strconv.FormatFloat(vv, 'f', -1, 64),
			float:      true,
			above64bit: false,
			number:     vv,
		}, nil
	case stringer, string:
		vv, err := numberFromString(convertStringerToString(v))
		if err != nil {
			return nil, IsNotNumber
		}
		return vv, nil
	default:
		return nil, IsNotNumber
	}
}
