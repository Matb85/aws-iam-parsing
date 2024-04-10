package types

/*
useful links:
- https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-role-policy.html
- https://docs.aws.amazon.com/IAM/latest/APIReference/API_PolicyDetail.html
- https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html
- https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_resource-based
- https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_delegate-permissions_examples.html
- https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html
*/

// Policy is a policy document.
type Policy struct {
	PolicyName     string          `json:"PolicyName" validate:"required"`
	PolicyDocument *PolicyDocument `json:"PolicyDocument" validate:"required"`
}

// Policy is a policy document.
type PolicyDocument struct {
	Id         *string           `json:"Id,omitempty"`
	Statements *StatementOrSlice `json:"Statement" validate:"required"`
	Version    *string           `json:"Version" validate:"required"`
}
