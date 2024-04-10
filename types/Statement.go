package types

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
