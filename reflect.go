package dt

import "reflect"

func MapReflectType(p reflect.Kind) Type {
	switch p {
	case reflect.Bool:
		return BoolType
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return NumberType
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return NumberType
	case reflect.String:
		return StringType
	default:
		return InvalidType
	}
}
