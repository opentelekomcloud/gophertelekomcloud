package users

// User contains the details of a RocketMQ instance user.
type User struct {
	// Username.
	AccessKey string `json:"access_key"`
	// Secret key.
	SecretKey string `json:"secret_key"`
	// IP address whitelist.
	WhiteRemoteAddress string `json:"white_remote_address"`
	// Whether the user is an administrator.
	Admin bool `json:"admin"`
	// Default topic permissions.
	//    PUB: publish permissions.
	//    SUB: subscribe permissions.
	//    PUB|SUB: subscribe and publish permissions.
	//    DENY: no permission.
	DefaultTopicPerm string `json:"default_topic_perm"`
	// Default consumer group permissions.
	//    SUB: subscribe permissions.
	//    DENY: no permission.
	DefaultGroupPerm string `json:"default_group_perm"`
	// Special topic permissions.
	TopicPerms []Permission `json:"topic_perms"`
	// Special consumer group permissions.
	GroupPerms []Permission `json:"group_perms"`
}

// Permission is a special permission of a user for a topic or a consumer group.
type Permission struct {
	// Topic or consumer group name.
	Name string `json:"name,omitempty"`
	// Topic permissions: PUB, SUB, PUB|SUB or DENY.
	// Consumer group permissions: SUB or DENY.
	Perm string `json:"perm,omitempty"`
}

// AccessPolicy is a user granted permissions for a topic or a consumer group.
type AccessPolicy struct {
	// Username.
	AccessKey string `json:"access_key"`
	// IP address whitelist.
	WhiteRemoteAddress string `json:"white_remote_address"`
	// Whether the user is an administrator.
	Admin bool `json:"admin"`
	// User permissions: PUB, SUB, PUB|SUB or DENY.
	Perm string `json:"perm"`
}

type ListAccessPoliciesOpts struct {
	// Offset where the query starts. Must be greater than or equal to 0.
	Offset int `q:"offset,omitempty"`
	// Number of records to query.
	Limit int `q:"limit,omitempty"`
}

// ListAccessPoliciesResponse is returned by ListTopicAccessPolicies
// and ListGroupAccessPolicies.
type ListAccessPoliciesResponse struct {
	// Users granted permissions.
	Policies []AccessPolicy `json:"policies"`
	// Total number of users.
	Total int `json:"total"`
	// Name of the topic or consumer group.
	Name string `json:"name"`
}
