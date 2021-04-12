package dt

type Type uint8

const (
	InvalidType Type = iota
	NilType
	BoolType
	NumberType
	StringType
)
