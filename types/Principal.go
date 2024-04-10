package types

// principal is a json-serializable type in a policy document.
type principal struct {
	AWS           *StringOrSlice `json:"AWS,omitempty"`
	CanonicalUser *StringOrSlice `json:"CanonicalUser,omitempty"`
	Federated     *StringOrSlice `json:"Federated,omitempty"`
	Service       *StringOrSlice `json:"Service,omitempty"`
}

// Principal is a Principal in a policy document.
type Principal struct {
	principal *principal
	str       string
}
