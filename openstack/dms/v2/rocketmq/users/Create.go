package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Username. Starts with a letter, consists of 7 to 64 characters, and contains
	// only letters, digits, hyphens (-), and underscores (_).
	AccessKey string `json:"access_key" required:"true"`
	// Secret key. 8 to 32 characters, containing at least three of the following
	// character types: uppercase letters, lowercase letters, digits and special
	// characters. Cannot be the username or the username spelled backwards.
	SecretKey string `json:"secret_key" required:"true"`
	// IP address whitelist, for example 192.168.1.*.
	WhiteRemoteAddress string `json:"white_remote_address,omitempty"`
	// Whether the user is an administrator.
	Admin *bool `json:"admin,omitempty"`
	// Default topic permissions: PUB, SUB, PUB|SUB or DENY.
	DefaultTopicPerm string `json:"default_topic_perm,omitempty"`
	// Default consumer group permissions: SUB or DENY.
	DefaultGroupPerm string `json:"default_group_perm,omitempty"`
	// Special topic permissions.
	TopicPerms []Permission `json:"topic_perms,omitempty"`
	// Special consumer group permissions.
	GroupPerms []Permission `json:"group_perms,omitempty"`
}

// Create a user of a RocketMQ instance.
// Send POST to /v2/{project_id}/instances/{instance_id}/users
func Create(client *golangsdk.ServiceClient, instanceID string, opts CreateOpts) (*User, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceID, "users").Build()
	if err != nil {
		return nil, err
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res User
	err = extract.Into(raw.Body, &res)
	return &res, err
}
