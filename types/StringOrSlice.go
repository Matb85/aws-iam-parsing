package types

import (
	"encoding/json"
	"errors"
	"matb85/remitly-home-assignment/utils"
)

// StringOrSlice is a type that can hold a string or a slice of strings.
// When unmarshalling JSON, it will preserve whether the original value was
// a singular string or a slice of strings.
type StringOrSlice struct {
	Values   []string
	Singular bool
}

func (s *StringOrSlice) UnmarshalJSON(data []byte) error {
	var tmp interface{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	slice, ok := tmp.([]interface{})
	if ok {
		values := []string{}
		for _, item := range slice {
			if _, ok := item.(string); !ok {
				return errors.New(utils.ErrorInvalidStringSlice)
			}
			values = append(values, item.(string))
			s.Singular = false
		}
		s.Values = values
		return nil
	}
	theString, ok := tmp.(string)
	if ok {
		s.Values = []string{theString}
		s.Singular = true
		return nil
	}
	return errors.New(utils.ErrorInvalidStringOrSlice)
}
