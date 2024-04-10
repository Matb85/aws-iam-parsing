package types

import (
	"matb85/remitly-home-assignment/utils"
	"testing"
)

func addToStringOrSlice(s *StringOrSlice, value ...string) {
	s.Values = append(s.Values, value...)
	if len(s.Values) != 1 {
		s.Singular = false
	}
}

func TestInvalidStringSliceJSON(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "NotSliceOfString",
			in:   `[{"foo": "bar"}]`,
			want: utils.ErrorInvalidStringSlice,
		},
		{
			name: "InvalidString",
			in:   `123`,
			want: utils.ErrorInvalidStringOrSlice,
		},
		{
			name: "InvalidJSON",
			in:   `{`,
			want: `unexpected end of JSON input`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var ss StringOrSlice
			err := ss.UnmarshalJSON([]byte(tc.in))
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
			if err.Error() != tc.want {
				t.Errorf("got '%s', want '%s'", err.Error(), tc.want)
			}
		})
	}
}

func TestStringOrSliceAdd(t *testing.T) {
	cases := []struct {
		name string
		in   *StringOrSlice
		add  []string
		want []string
	}{

		{
			name: "Singular",
			in:   newStringOrSlice(true, "arn:aws:iam::123456789012:root"),
			add:  []string{"arn:aws:iam::123456789012:root"},
			want: []string{"arn:aws:iam::123456789012:root", "arn:aws:iam::123456789012:root"},
		},
		{
			name: "Empty",
			in:   newStringOrSlice(false, "arn:aws:iam::123456789012:root"),
			add:  []string{},
			want: []string{"arn:aws:iam::123456789012:root"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			addToStringOrSlice(tc.in, tc.add...)
			if len(tc.in.Values) != len(tc.want) {
				t.Fatalf("got '%d', want '%d'", len(tc.in.Values), len(tc.want))
			}
			for i, v := range tc.in.Values {
				if v != tc.want[i] {
					t.Fatalf("got '%s', want '%s'", v, tc.want[i])
				}
			}
		})
	}
}
