package dt

import "math"

var FloatEps = 1e-8 //容忍度

func FloatEqual(a, b float64) bool {
	return FloatEqualC(a, b, FloatEps)
}

func FloatEqualC(a, b, eps float64) bool {
	if math.Abs(a-b) < eps {
		return true
	}
	return false
}

func FloatIsInt(f float64) bool {
	return FloatIsIntC(f, FloatEps)
}

func FloatIsInt64(f float64) bool {
	if f > math.MaxInt64 || f < math.MinInt64 {
		return false
	}

	return FloatIsInt(f)
}

func FloatToInt64(f float64) int64 {
	if f > math.MaxInt64 {
		return math.MaxInt64
	}
	if f < math.MinInt64 {
		return math.MinInt64
	}

	return int64(f)
}

func FloatIsIntC(f, eps float64) bool {
	return FloatEqualC(math.Mod(f, 1.0), 0, eps)
}
