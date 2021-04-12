package dt

import (
	"fmt"
	"github.com/ImSingee/tt"
	"testing"
)

func TestParseBool(t *testing.T) {
	// 非 0 数字均为 true
	trueCases := []interface{}{
		true,
		"1", "ok", "001",
		1, -1, 1.1,
	}
	falseCases := []interface{}{
		false,
		"0", "false", "000",
		0, 0.0,
	}

	for _, c := range trueCases {
		t.Run(fmt.Sprintf("Should Be True: %#+v", c), func(t *testing.T) {
			b, ok := ParseBool(c)
			tt.AssertTrue(t, ok)
			tt.AssertTrue(t, b)
		})
	}

	for _, c := range falseCases {
		t.Run(fmt.Sprintf("Should Be False: %#+v", c), func(t *testing.T) {
			b, ok := ParseBool(c)
			tt.AssertTrue(t, ok)
			tt.AssertFalse(t, b)
		})
	}

}
