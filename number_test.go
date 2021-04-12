package dt

import (
	"fmt"
	"github.com/ImSingee/tt"
	"testing"
)

func TestNewGenericNumber(t *testing.T) {
	// 整数 int/int8/int16/int32/int64
	number, err := newGenericNumber(int(123))
	tt.AssertIsNil(t, err)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertEqual(t, int64(123), number.number)

	number, err = newGenericNumber(int8(123))
	tt.AssertIsNil(t, err)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertEqual(t, int64(123), number.number)

	number, err = newGenericNumber(int16(123))
	tt.AssertIsNil(t, err)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertEqual(t, int64(123), number.number)

	number, err = newGenericNumber(int32(123))
	tt.AssertIsNil(t, err)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertEqual(t, int64(123), number.number)

	number, err = newGenericNumber(int64(123))
	tt.AssertIsNil(t, err)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertEqual(t, int64(123), number.number)

	// 正整数 uint/uint8/uint16/uint32/uint64
	number, err = newGenericNumber(uint(123))
	tt.AssertIsNil(t, err)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertEqual(t, uint64(123), number.number)

	number, err = newGenericNumber(uint8(123))
	tt.AssertIsNil(t, err)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertEqual(t, uint64(123), number.number)

	number, err = newGenericNumber(uint16(123))
	tt.AssertIsNil(t, err)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertEqual(t, uint64(123), number.number)

	number, err = newGenericNumber(uint32(123))
	tt.AssertIsNil(t, err)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertEqual(t, uint64(123), number.number)

	number, err = newGenericNumber(uint64(123))
	tt.AssertIsNil(t, err)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertEqual(t, uint64(123), number.number)

	// 浮点数 float32/float64
	number, err = newGenericNumber(float32(123.45))
	tt.AssertIsNil(t, err)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	//tt.AssertEqual(t, float32(123.45), number.number)

	number, err = newGenericNumber(123.45)
	tt.AssertIsNil(t, err)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertEqual(t, 123.45, number.number)

	// 字符串
	number, err = newGenericNumber("123")
	tt.AssertIsNil(t, err)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertEqual(t, int64(123), number.number)

	number, err = newGenericNumber("123.456")
	tt.AssertIsNil(t, err)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertEqual(t, 123.456, number.number)
}
