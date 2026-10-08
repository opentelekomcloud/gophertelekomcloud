package gpfs

// ListFSInput is the input parameter of ListFS function
type ListFSInput struct {
	BucketType string
}

// ListFSOutput is the result of ListFS function
type ListFSOutput struct {
	BaseModel
	Owner   Owner    `xml:"Owner"`
	Buckets []Bucket `xml:"Buckets>Bucket"`
}

// CreateFSInput is the input parameter of CreateBucket function
type CreateFSInput struct {
	BucketLocation
	FSName     string `xml:"-"`
	Redundancy string `xml:"-"`
	BucketType string `xml:"-"`
}

// AccessRuleCondition defines the clients to which a file system access rule applies.
type AccessRuleCondition struct {
	// SourceVPC is the ID of the VPC allowed to access the file system.
	SourceVPC string `json:"SourceVpc"`
	// VPCSourceIP optionally lists IP addresses or CIDR ranges in the source VPC.
	// The service currently ignores this field.
	VPCSourceIP []string `json:"VpcSourceIp,omitempty"`
}

// AccessRule defines access to a general-purpose file system from a VPC.
type AccessRule struct {
	// ID is an optional statement identifier.
	ID string `json:"Sid,omitempty"`
	// Action is FullControl for read/write access or Read for read-only access.
	Action AccessRuleAction `json:"Action"`
	// Effect must be Allow.
	Effect AccessRuleEffect `json:"Effect"`
	// Condition identifies the VPC and, optionally, its allowed source addresses.
	Condition AccessRuleCondition `json:"Condition"`
}

// CreateFSAccessRulesInput is the request body for configuring file system access rules.
type CreateFSAccessRulesInput struct {
	// FSName is the name of the file system.
	FSName string `json:"-"`
	// Statement contains the access rules to configure.
	Statement []AccessRule `json:"Statement"`
}

// GetFSAccessRulesOutput is the result of querying file system access rules.
type GetFSAccessRulesOutput struct {
	BaseModel
	Statement []AccessRule `json:"Statement"`
}
