package types

// ConditionValue is a type that can hold an indivual or slice of string, bool or float64.
// When unarshalling JSON, it will preserve whether the original value was
// singular or a slice.
type ConditionValue struct {
	StrValues  []string
	BoolValues []bool
	NumValues  []float64
	Singular   bool
}
