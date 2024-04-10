package types

import (
	"encoding/json"
	"errors"
	"matb85/remitly-home-assignment/utils"
)

// ConditionValue is a type that can hold an indivual or slice of string, bool or float64.
// When unarshalling JSON, it will preserve whether the original value was
// singular or a slice.
type ConditionValue struct {
	StrValues  []string
	BoolValues []bool
	NumValues  []float64
	Singular   bool
}

func (c *ConditionValue) UnmarshalJSON(data []byte) error {
	var tmp interface{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	theString, ok := tmp.(string)
	if ok {
		c.StrValues = []string{theString}
		c.Singular = true
		return nil
	}
	theBool, ok := tmp.(bool)
	if ok {
		c.BoolValues = []bool{theBool}
		c.Singular = true
		return nil
	}
	theFloat, ok := tmp.(float64)
	if ok {
		c.NumValues = []float64{theFloat}
		c.Singular = true
		return nil
	}

	slice, ok := tmp.([]interface{})
	if ok {
		strValues := []string{}
		boolValues := []bool{}
		numValues := []float64{}
		for _, item := range slice {
			switch item.(type) {
			case string:
				strValues = append(strValues, item.(string))
			case bool:
				boolValues = append(boolValues, item.(bool))
			case float64: // all numbers are float64
				numValues = append(numValues, item.(float64))
			default:
				return errors.New(utils.ErrorInvalidConditionValueSlice)
			}
			c.Singular = false
		}
		c.StrValues = strValues
		c.BoolValues = boolValues
		c.NumValues = numValues
		return nil
	}

	return errors.New(utils.ErrorInvalidConditionValue)
}
