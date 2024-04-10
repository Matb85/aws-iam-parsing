package types

// StringOrSlice is a type that can hold a string or a slice of strings.
// When unarshalling JSON, it will preserve whether the original value was
// a singular string or a slice of strings.
type StringOrSlice struct {
	Values   []string
	Singular bool
}
