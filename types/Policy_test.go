package types

import (
	"bytes"
	"encoding/json"
	"matb85/remitly-home-assignment/utils"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// newStringOrSlice creates a new StringOrSlice. If singular is true and
// there is only one element, the structure will be marshaled as a string
// instead of a slice.
func newStringOrSlice(singular bool, values ...string) *StringOrSlice {
	return &StringOrSlice{
		Values:   values,
		Singular: singular,
	}
}

// newAWSPrincipal creates a new Principal that matches an AWS account.
func newAWSPrincipal(aws ...string) *Principal {
	return &Principal{
		principal: &principal{
			AWS: newStringOrSlice(true, aws...),
		},
	}
}

func TestDisallowUnknownFields(t *testing.T) {
	cases := []struct {
		name    string
		in      string
		wantErr string
	}{
		{
			name: "AllowUnknownFieldsInPolicy",
			in: `{
				"Version": "2012-10-17",
				"NewField": "NewValue",
				"Statement": [
					{
						"Effect": "Allow",
						"Action": "s3:GetObject",
						"Resource": "arn:aws:s3:::my_corporate_bucket/exampleobject.png"
					}
				]
			}`,
			wantErr: `json: unknown field "NewField"`,
		},
		{
			name: "AllowUnknownFieldsInStatement",
			in: `{
				"Version": "2012-10-17",
				"Statement": [
					{
						"Effect": "Allow",
						"Action": "s3:GetObject",
						"Resource": "arn:aws:s3:::my_corporate_bucket/exampleobject.png",
						"NewField": "NewValue"
					}
				]
			}`,
			wantErr: `StatementOrSlice is not a slice of statements: json: unknown field "NewField"`,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var p PolicyDocument
			decoder := json.NewDecoder(bytes.NewBufferString(tc.in))
			decoder.DisallowUnknownFields()
			err := decoder.Decode(&p)
			if err == nil {
				t.Fatalf("expect error, got none")
			}
			if err.Error() != tc.wantErr {
				t.Fatalf("expect error %q, got %q", tc.wantErr, err)
			}
		})
	}
}

func TestStatementOrSliceUnmarshalJSON(t *testing.T) {
	cases := []struct {
		name         string
		in           string
		want         []Statement
		wantSingular bool
		wantErr      string
	}{
		{
			name: "SingleStatement",
			in: `{
				"Effect": "Allow",
				"Action": "s3:GetObject",
				"Resource": "arn:aws:s3:::my_corporate_bucket/exampleobject.png",
				"Principal": {
					"AWS": "123456789012"
				}
			}`,
			want: []Statement{
				{
					Effect:    utils.EffectAllow,
					Action:    newStringOrSlice(true, "s3:GetObject"),
					Resource:  newStringOrSlice(true, "arn:aws:s3:::my_corporate_bucket/exampleobject.png"),
					Principal: newAWSPrincipal("123456789012"),
				},
			},
			wantSingular: true,
		},
		{
			name: "SliceStatement",
			in: `[
				{
					"Effect": "Allow",
					"Action": "s3:GetObject",
					"Resource": "arn:aws:s3:::my_corporate_bucket/exampleobject.png",
					"Principal": {
						"AWS": "123456789012"
					}
				}
			]`,
			want: []Statement{
				{
					Effect:    utils.EffectAllow,
					Action:    newStringOrSlice(true, "s3:GetObject"),
					Resource:  newStringOrSlice(true, "arn:aws:s3:::my_corporate_bucket/exampleobject.png"),
					Principal: newAWSPrincipal("123456789012"),
				},
			},
			wantSingular: false,
		},
		{
			name: "InvalidJSON",
			in: `{
				"Effect": "Allow",
				"Action": "s3:GetObject",
				`,
			wantErr:      "unexpected end of JSON input",
			wantSingular: false,
		},
		{
			name:    "BooleanJSON",
			in:      `true`,
			wantErr: utils.ErrorInvalidStatementOrSlice,
		},
		{
			name:    "BadJSON",
			in:      `{`,
			wantErr: `unexpected end of JSON input`,
		},
		{
			name: "InvalidList",
			in: `[
				{
					"Effect": "Allow",
					"NotAField": "s3:GetObject"
				}
			]`,
			wantErr:      `StatementOrSlice is not a slice of statements: json: unknown field "NotAField"`,
			wantSingular: false,
		},
		{
			name: "InvalidStatement",
			in: `{
				"Effect": "Allow",
				"NotAField": "s3:GetObject"
			}`,
			wantErr:      `StatementOrSlice must be a single Statement or a slice of Statements: json: unknown field "NotAField"`,
			wantSingular: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var s StatementOrSlice
			err := s.UnmarshalJSON([]byte(tc.in))
			if err != nil {
				if tc.wantErr == "" {
					t.Fatalf("expect no error, got %v", err)
				}
				if err.Error() != tc.wantErr {
					t.Fatalf("expect error %q, got %q", tc.wantErr, err)
				}
				return
			}
			if len(tc.want) != len(s.Values) {
				t.Errorf("got '%d', want '%d'", len(s.Values), len(tc.want))
				return
			}
			if !cmp.Equal(tc.want, s.Values, cmpopts.IgnoreUnexported(StringOrSlice{}, Principal{})) {
				t.Errorf("%s", cmp.Diff(tc.want, s.Values, cmpopts.IgnoreUnexported(StringOrSlice{}, Principal{})))
				return
			}
			if tc.wantSingular != s.Singular {
				t.Errorf("got '%t', want '%t'", s.Singular, tc.wantSingular)
			}
		})
	}
}
