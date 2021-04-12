package dt

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

type Number interface {
	AsNumber() *GenericNumber
}

type GenericNumber struct {
	literal string

	above64bit bool // 是否不足以使用 64 位表示 (number 类型为 *big.Int / *big.Float)
	float      bool // 是否为浮点数（number 类型为 float64 / *big.Float）
	unsigned   bool // 是否有符号（number 类型为 uint64）

	number interface{} // 类型为 int64/uint64/float64/*big.Int/*big.Float
}

func ConvertSignedIntToInt64(num interface{}) (int64, bool) {
	switch v := num.(type) {
	case int:
		return int64(v), true
	case int8:
		return int64(v), true
	case int16:
		return int64(v), true
	case int32:
		return int64(v), true
	case int64:
		return v, true
	default:
		return 0, false
	}
}

func ConvertUnsignedIntToUInt64(num interface{}) (uint64, bool) {
	switch v := num.(type) {
	case uint:
		return uint64(v), true
	case uint8:
		return uint64(v), true
	case uint16:
		return uint64(v), true
	case uint32:
		return uint64(v), true
	case uint64:
		return v, true
	default:
		return 0, false
	}
}

func ConvertFloatToFloat64(num interface{}) (float64, bool) {
	switch v := num.(type) {
	case float32:
		return float64(v), true
	case float64:
		return v, true
	default:
		return 0, false
	}
}

// 从字符串中解析整数
func IntFromString(num string) (*GenericNumber, bool) {
	v, err := strconv.ParseInt(num, 10, 64)
	if err == nil {
		return &GenericNumber{
			literal:    num,
			float:      false,
			unsigned:   false,
			above64bit: false,
			number:     v,
		}, true
	}
	return nil, false
}

func UIntFromString(num string) (*GenericNumber, bool) {
	vv, err := strconv.ParseUint(num, 10, 64)
	if err == nil {
		return &GenericNumber{
			literal:    num,
			float:      false,
			unsigned:   true,
			above64bit: false,
			number:     vv,
		}, true
	}
	return nil, false
}

func BigIntFromString(num string) (*GenericNumber, bool) {
	vvv, ok := new(big.Int).SetString(num, 10)
	if ok {
		return &GenericNumber{
			literal:    num,
			above64bit: true,
			float:      false,
			number:     vvv,
		}, true
	}
	return nil, false
}

func FloatFromString(num string) (*GenericNumber, bool) {
	v, err := strconv.ParseFloat(num, 64)
	if err == nil {
		return &GenericNumber{
			literal:    num,
			float:      true,
			unsigned:   false,
			above64bit: false,
			number:     v,
		}, true
	}
	return nil, false
}

func BigFloatFromString(num string) (*GenericNumber, bool) {
	f, ok := new(big.Float).SetString(num)
	if ok {
		return &GenericNumber{
			literal:    num,
			above64bit: true,
			float:      true,
			number:     f,
		}, true
	}
	return nil, false
}

// 从字符串获取对应表示的数字值
func NumberFromString(num string) (*GenericNumber, bool) {
	// 看看有没有点（是整数还是小数）
	if !strings.Contains(num, ".") {
		// 尝试作为整数解析
		if v, ok := IntFromString(num); ok {
			return v, true
		}
		// 尝试作为正整数解析
		if v, ok := UIntFromString(num); ok {
			return v, true
		}
		// 尝试大数解析
		if v, ok := BigIntFromString(num); ok {
			return v, true
		}
	} else {
		// 尝试作为小数解析
		if v, ok := FloatFromString(num); ok {
			return v, true
		}

		//// 尝试大数解析
		//v, err = BigFloatFromString(num)
		//if err == nil {
		//	return v, nil
		//}
	}

	// 返回错误 （不是数字）
	return nil, false
}

func NumberFromBasicInt(num interface{}) (*GenericNumber, bool) {
	if vv, ok := ConvertSignedIntToInt64(num); ok {
		return &GenericNumber{
			literal:    strconv.FormatInt(vv, 10),
			float:      false,
			unsigned:   false,
			above64bit: false,
			number:     vv,
		}, true
	}

	return nil, false
}

func NumberFromBasicUInt(num interface{}) (*GenericNumber, bool) {
	if vv, ok := ConvertUnsignedIntToUInt64(num); ok {

		return &GenericNumber{
			literal:    strconv.FormatUint(vv, 10),
			float:      false,
			unsigned:   true,
			above64bit: false,
			number:     vv,
		}, true
	}
	return nil, false
}

func NumberFromBasicFloat(num interface{}) (*GenericNumber, bool) {
	if vv, ok := ConvertFloatToFloat64(num); ok {
		return &GenericNumber{
			literal:    ToString(vv),
			float:      true,
			above64bit: false,
			number:     vv,
		}, true
	}

	return nil, false
}

func NumberFromBasicType(num interface{}) (*GenericNumber, bool) {
	if v, ok := NumberFromBasicInt(num); ok {
		return v, true
	}

	if v, ok := NumberFromBasicUInt(num); ok {
		return v, true
	}

	if v, ok := NumberFromBasicFloat(num); ok {
		return v, true
	}

	return nil, false
}

func ParseNumber(num interface{}) (*GenericNumber, bool) {
	if v, ok := num.(*GenericNumber); ok {
		return v, true
	}
	if v, ok := NumberFromBasicType(num); ok {
		return v, true
	}

	switch v := num.(type) {
	case Stringer, string:
		if vv, ok := NumberFromString(ToString(v)); ok {
			return vv, true
		}

		return nil, false
	}

	return nil, false
}

// 查看是否可以将值转换成 int64
func (num *GenericNumber) IsInt64() bool {
	if !num.float {
		if num.above64bit {
			return num.number.(*big.Int).IsInt64()
		}

		if num.unsigned {
			return num.number.(uint64) <= math.MaxInt64
		} else {
			return true
		}
	} else {
		if num.above64bit {
			return num.number.(*big.Float).IsInt()
		} else {
			return FloatIsInt64(num.number.(float64))
		}
	}
}

// 将值转换为 int64
// 如果 IsInt64 == false，则返回结果是不确定的
func (num *GenericNumber) Int64() int64 {
	if num.float {
		return FloatToInt64(num.number.(float64))
	}
	if num.above64bit {
		return num.number.(*big.Int).Int64()
	}
	if num.unsigned {
		return int64(num.number.(uint64))
	} else {
		return num.number.(int64)
	}
}

func (num *GenericNumber) String() string {
	return num.literal
}
