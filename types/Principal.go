package types

import "encoding/json"

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

func (p *Principal) UnmarshalJSON(data []byte) error {
	var tmp interface{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	str, ok := tmp.(string)
	if ok {
		p.str = str
		return nil
	}
	principal := &principal{}
	err = json.Unmarshal(data, principal)
	if err != nil {
		return err
	}
	p.principal = principal
	return nil
}
