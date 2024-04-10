package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"matb85/remitly-home-assignment/utils"
)

// Statement is a single statement in a policy document.
type Statement struct {
	Action       *StringOrSlice                        `json:"Action,omitempty" validate:"required"`
	Condition    map[string]map[string]*ConditionValue `json:"Condition,omitempty"`
	Effect       string                                `json:"Effect" validate:"required"`
	NotAction    *StringOrSlice                        `json:"NotAction,omitempty"`
	NotResource  *StringOrSlice                        `json:"NotResource,omitempty"`
	Principal    *Principal                            `json:"Principal,omitempty"`
	NotPrincipal *Principal                            `json:"NotPrincipal,omitempty"`
	Resource     *StringOrSlice                        `json:"Resource,omitempty" validate:"required"`
	Sid          string                                `json:"Sid,omitempty"`
}

// StatementOrSlice represents Statements that can be marshaled to a single Statement or a slice of Statements.
type StatementOrSlice struct {
	Values   []Statement
	Singular bool
}

func (s *StatementOrSlice) UnmarshalJSON(data []byte) error {
	var tmp interface{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	// verify whether it is a single object or an array
	_, ok := tmp.([]interface{})
	if ok {
		values := []Statement{}
		decoder := json.NewDecoder(bytes.NewReader(data))
		decoder.DisallowUnknownFields()
		err = decoder.Decode(&values)
		if err != nil {
			return fmt.Errorf("%s: %v", utils.ErrorInvalidStatementSlice, err)

		}
		s.Values = values
		s.Singular = false
		return nil
	}
	_, ok = tmp.(map[string]interface{})
	if ok {
		value := Statement{}
		decoder := json.NewDecoder(bytes.NewReader(data))
		decoder.DisallowUnknownFields()
		err = decoder.Decode(&value)
		if err != nil {
			return fmt.Errorf("%s: %v", utils.ErrorInvalidStatementOrSlice, err)
		}
		s.Values = []Statement{value}
		s.Singular = true
		return nil
	}
	return errors.New(utils.ErrorInvalidStatementOrSlice)
}
