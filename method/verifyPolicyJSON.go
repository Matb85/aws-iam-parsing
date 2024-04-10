package method

import (
	"bytes"
	"encoding/json"
	"fmt"
	"matb85/remitly-home-assignment/types"

	"github.com/go-playground/validator/v10"
)

func VerifyPolicyJSON(data []byte) bool {
	//fmt.Println(string(data))
	hasTopLevelElements := true
	var p types.Policy
	decoder := json.NewDecoder(bytes.NewBuffer(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&p)
	if err != nil {
		fmt.Println("JSON error -", err)
		// https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_resource-based
		// The policyName & PolicyDocument are optional, try to parse json without them
		decoder = json.NewDecoder(bytes.NewBuffer(data))
		decoder.DisallowUnknownFields()
		err = decoder.Decode(&p.PolicyDocument)
		if err != nil {
			fmt.Println("JSON error -", err)
			return false
		}

		hasTopLevelElements = false
	}

	// Create a new validator instance
	validate := validator.New()
	// Validate the Policy or PolicyDocument struct
	if hasTopLevelElements {
		err = validate.Struct(p)
	} else {
		err = validate.Struct(p.PolicyDocument)
	}
	// try again only on p.PolicyDocument
	if err != nil {
		// Validation failed, handle the error
		fmt.Println(err)
		// https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_resource-based
		decoder = json.NewDecoder(bytes.NewBuffer(data))
		decoder.DisallowUnknownFields()
		err = decoder.Decode(&p.PolicyDocument)
		if err != nil {
			fmt.Println(err)
			return false
		}

		// Validate the PolicyDocument struct
		err = validate.Struct(p.PolicyDocument)
		if err != nil {
			fmt.Println(err)
			return false
		}

	}

	var statements []types.Statement = p.PolicyDocument.Statements.Values
	containsAsterisk := false
	for _, s := range statements {
		// Validate the Statement struct
		err = validate.Struct(s)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, val := range s.Resource.Values {
			if val == "*" {
				if !containsAsterisk {
					containsAsterisk = true
				} else {
					// if there are more asterisks, return false
					return false
				}
			}
		}

		// fmt.Println(index, s.Resource.Values)
	}

	return containsAsterisk
}
