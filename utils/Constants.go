package utils

const (
	ErrorInvalidConditionValueSlice = "field not slice of string, bool or float64"
	ErrorInvalidConditionValue      = "field neither slice of string, bool, or float64 or string, bool or float64"

	// See https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_effect.html
	EffectAllow = "Allow"
	EffectDeny  = "Deny"

	// See https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_version.html
	Version2012_10_17 = "2012-10-17"
	Version2008_10_17 = "2008-10-17"
	VersionLatest     = Version2012_10_17

	ErrorInvalidStatementSlice   = "StatementOrSlice is not a slice of statements"
	ErrorInvalidStatementOrSlice = "StatementOrSlice must be a single Statement or a slice of Statements"

	ErrorInvalidStringOrSlice = "field neither slice of string or string"
	ErrorInvalidStringSlice   = "field not slice of string"

	ErrorPolicyMissingField = "PolicyDocument is missing the following field: "
)
