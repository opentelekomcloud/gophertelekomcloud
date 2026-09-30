package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Secret key. The API rejects an empty key, so the current key
	// has to be passed to keep it unchanged.
	SecretKey string `json:"secret_key" required:"true"`
	// IP address whitelist, for example 192.168.1.*. An empty value clears it.
	WhiteRemoteAddress *string `json:"white_remote_address,omitempty"`
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

// Update the parameters of a specified user of a RocketMQ instance.
// Send PUT to /v2/{project_id}/instances/{instance_id}/users/{user_name}
func Update(client *golangsdk.ServiceClient, instanceID, userName string, opts UpdateOpts) (*User, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceID, "users", userName).Build()
	if err != nil {
		return nil, err
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res User
	err = extract.Into(raw.Body, &res)
	return &res, err
}
