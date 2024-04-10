package method

import (
	"testing"

	"encoding/json"
	"fmt"
	"os"
)

func TestFalseCorrectPolicies(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	fixture := path + "/../test/false_correct_policies.json"

	reader, err := os.ReadFile(fixture)
	if err != nil {
		t.Fatal(err)
	}

	policies := []interface{}{}
	err = json.Unmarshal(reader, &policies)
	if err != nil {
		t.Fatal(err)
	}

	for i, policyIface := range policies {
		t.Run(fmt.Sprintf("Unmarshal valid policies that should return false %d", i), func(t *testing.T) {
			// Yes we re-marshal the policy here, but this is the only way to
			// get the json package to use the MarshalJSON() method on each
			// Principal struct individually.
			b, err := json.MarshalIndent(policyIface, "", "  ")
			if err != nil {
				t.Fatal(err)
			}
			value := VerifyPolicyJSON(b)
			if value != false {
				testName := struct {
					Id string `json:"Id"`
				}{}
				uerr := json.Unmarshal(b, &testName)
				if uerr != nil {
					t.Fatal(uerr)
				}
				t.Fatalf("Error testing policy %s: %v", testName.Id, err)
			}
		})
	}
}
func TestFalseInvalidPolicies(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(path)
	fixture := path + "/../test/false_invalid_policies.json"

	reader, err := os.ReadFile(fixture)
	if err != nil {
		t.Fatal(err)
	}

	policies := []interface{}{}
	err = json.Unmarshal(reader, &policies)
	if err != nil {
		t.Fatal(err)
	}

	for i, policyIface := range policies {
		t.Run(fmt.Sprintf("Unmarshal valid policies that should return false %d", i), func(t *testing.T) {
			// Yes we re-marshal the policy here, but this is the only way to
			// get the json package to use the MarshalJSON() method on each
			// Principal struct individually.
			b, err := json.MarshalIndent(policyIface, "", "  ")
			if err != nil {
				t.Fatal(err)
			}
			value := VerifyPolicyJSON(b)
			if value != false {
				testName := struct {
					Id string `json:"Id"`
				}{}
				uerr := json.Unmarshal(b, &testName)
				if uerr != nil {
					t.Fatal(uerr)
				}
				t.Fatalf("Error testing policy %s: %v", testName.Id, err)
			}
		})
	}
}
func TestTrueCorrectPolicies(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	fixture := path + "/../test/true_correct_policies.json"

	reader, err := os.ReadFile(fixture)
	if err != nil {
		t.Fatal(err)
	}

	policies := []interface{}{}
	err = json.Unmarshal(reader, &policies)
	if err != nil {
		t.Fatal(err)
	}

	for i, policyIface := range policies {
		t.Run(fmt.Sprintf("Unmarshal valid policies that should return true %d", i), func(t *testing.T) {
			// Yes we re-marshal the policy here, but this is the only way to
			// get the json package to use the MarshalJSON() method on each
			// Principal struct individually.
			b, err := json.MarshalIndent(policyIface, "", "  ")
			if err != nil {
				t.Fatal(err)
			}
			value := VerifyPolicyJSON(b)
			if value != true {
				testName := struct {
					Id string `json:"Id"`
				}{}
				uerr := json.Unmarshal(b, &testName)
				if uerr != nil {
					t.Fatal(uerr)
				}
				t.Fatalf("Error testing policy %s: %v", testName.Id, err)
			}
		})
	}
}
