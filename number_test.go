package dt

import (
	"fmt"
	"github.com/ImSingee/tt"
	"math/big"
	"testing"
)

func TestNumberFromString(t *testing.T) {
	// 整数 int64
	// 范围 [-1 << 63, 1<<63 - 1] => [-9223372036854775808, 9223372036854775807]
	//fmt.Printf("int64 range: [%d, %d]\n", math.MinInt64, math.MaxInt64)
	number, ok := NumberFromString("123")
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertFalse(t, number.float)
	tt.AssertFalse(t, number.unsigned)
	tt.AssertFalse(t, number.above64bit)
	tt.AssertEqual(t, int64(123), number.number)

	// 正整数 uint64
	// 范围 [0, 1<<64 - 1] => [0, 18446744073709551615]
	//fmt.Printf("uint64 range: [0, %d]\n", uint64(math.MaxUint64))
	number, ok = NumberFromString("9223372036854775999") // 大于 MaxInt64 的数字应当被推导为 uint64
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertFalse(t, number.float)
	tt.AssertTrue(t, number.unsigned)
	tt.AssertFalse(t, number.above64bit)
	tt.AssertEqual(t, uint64(9223372036854775999), number.number)

	// 小数
	number, ok = NumberFromString("123.456")
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertTrue(t, number.float)
	tt.AssertFalse(t, number.above64bit)
	tt.AssertEqual(t, 123.456, number.number)

	number, ok = NumberFromString("18446744073709559999.888")
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertTrue(t, number.float)
	tt.AssertFalse(t, number.above64bit)
	tt.AssertEqual(t, 18446744073709559999.888, number.number)

	// 大整数
	number, ok = NumberFromString("18446744073709559999")
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertFalse(t, number.float)
	tt.AssertTrue(t, number.above64bit)
	tt.AssertFalse(t, number.number.(*big.Int).IsUint64())
	tt.AssertFalse(t, number.number.(*big.Int).IsInt64())
	// python: list((18446744073709559999).to_bytes(9, 'big'))
	// [1, 0, 0, 0, 0, 0, 0, 32, 191]
	tt.AssertEqual(t, []byte{1, 0, 0, 0, 0, 0, 0, 32, 191}, number.number.(*big.Int).Bytes())

	// 大小数
	// MaxFloat64 = 1.797693134862315708145274237317043567981e+308 // 2**1023 * (2**53 - 1) / 2**52
	//s := "18567834567832456789234567899765433567897656783456783245678923456789976543356789765678345678324567892345678997654335678976567834567832456789234567899765433567897656783456783245678923456789976543356789765678345678324567892345678997654335678976567834567832456789234567899765433567897656783456783245678923456789976543356789765678345678324567892345678997654335678976567834567832456789234567899765433567897656783456783245678923456789976543356789765678345678324567892345678997654335678976567834567832456789234567899765433567897656783456783245678923456789976543356789765678345678324567892345678997654335678976567834567832456789234567899765433567897602456783456783245678923456789976543356789768927459823459872905472980534727864781239482345678987654234567376489231486432764178919265478429363896785962378139468721463178463827136478164217834617349817234891742839749812374981473891423798347198234719234238947834658924374798234446744073709559999.123999"
	//number, ok = NumberFromString(s)
	//tt.AssertTrue(t, ok)
	//fmt.Printf("%#+v (%T)\n", number, number.number)
	//tt.AssertTrue(t, number.float)
	//tt.AssertTrue(t, number.above64bit)
	////tt.AssertFalse(t, number.number.(*big.Float).IsInt())
	//tt.AssertFalse(t, number.number.(*big.Float).IsInf())
	//tt.AssertEqual(t, s, fmt.Sprintf("%f", number.number.(*big.Float)))
}

func TestNewGenericNumber(t *testing.T) {
	// 整数 int/int8/int16/int32/int64
	number, ok := ParseNumber(int(123))
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertFalse(t, number.float)
	tt.AssertFalse(t, number.unsigned)
	tt.AssertFalse(t, number.above64bit)
	tt.AssertEqual(t, int64(123), number.number)

	number, ok = ParseNumber(int8(123))
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertFalse(t, number.float)
	tt.AssertFalse(t, number.unsigned)
	tt.AssertFalse(t, number.above64bit)
	tt.AssertEqual(t, int64(123), number.number)

	number, ok = ParseNumber(int16(123))
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertFalse(t, number.float)
	tt.AssertFalse(t, number.unsigned)
	tt.AssertFalse(t, number.above64bit)
	tt.AssertEqual(t, int64(123), number.number)

	number, ok = ParseNumber(int32(123))
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertFalse(t, number.float)
	tt.AssertFalse(t, number.unsigned)
	tt.AssertFalse(t, number.above64bit)
	tt.AssertEqual(t, int64(123), number.number)

	number, ok = ParseNumber(int64(123))
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertFalse(t, number.float)
	tt.AssertFalse(t, number.unsigned)
	tt.AssertFalse(t, number.above64bit)
	tt.AssertEqual(t, int64(123), number.number)

	// 正整数 uint/uint8/uint16/uint32/uint64
	number, ok = ParseNumber(uint(123))
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertFalse(t, number.float)
	tt.AssertTrue(t, number.unsigned)
	tt.AssertFalse(t, number.above64bit)
	tt.AssertEqual(t, uint64(123), number.number)

	number, ok = ParseNumber(uint8(123))
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertFalse(t, number.float)
	tt.AssertTrue(t, number.unsigned)
	tt.AssertFalse(t, number.above64bit)
	tt.AssertEqual(t, uint64(123), number.number)

	number, ok = ParseNumber(uint16(123))
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertFalse(t, number.float)
	tt.AssertTrue(t, number.unsigned)
	tt.AssertFalse(t, number.above64bit)
	tt.AssertEqual(t, uint64(123), number.number)

	number, ok = ParseNumber(uint32(123))
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertFalse(t, number.float)
	tt.AssertTrue(t, number.unsigned)
	tt.AssertFalse(t, number.above64bit)
	tt.AssertEqual(t, uint64(123), number.number)

	number, ok = ParseNumber(uint64(123))
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertFalse(t, number.float)
	tt.AssertTrue(t, number.unsigned)
	tt.AssertFalse(t, number.above64bit)
	tt.AssertEqual(t, uint64(123), number.number)

	// 浮点数 float32/float64
	number, ok = ParseNumber(float32(123.45))
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertTrue(t, number.float)
	tt.AssertFalse(t, number.above64bit)
	//tt.AssertEqual(t, float32(123.45), number.number)

	number, ok = ParseNumber(123.45)
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertTrue(t, number.float)
	tt.AssertFalse(t, number.above64bit)
	tt.AssertEqual(t, 123.45, number.number)

	// 字符串
	number, ok = ParseNumber("123")
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertEqual(t, int64(123), number.number)

	number, ok = ParseNumber("123.456")
	tt.AssertTrue(t, ok)
	fmt.Printf("%#+v (%T)\n", number, number.number)
	tt.AssertEqual(t, 123.456, number.number)

	// 异常类别
	number, ok = ParseNumber(nil)
	tt.AssertFalse(t, ok)
}
