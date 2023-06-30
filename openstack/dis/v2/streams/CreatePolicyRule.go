package streams

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreatePolicyRuleOpts struct {
	// Name of the stream for which you want to add an authorization policy.
	// Maximum: 64
	StreamName string `json:"stream_name" required:"true"`
	// Unique ID of the stream.
	StreamId string `json:"stream_id"`
	// Authorized users.
	// If the permission is granted to a specified tenant, the format is domainName.*.
	// If the permission is granted to a specified sub-user of a tenant, the format is domainName.userName.
	// Multiple accounts can be added and separated by commas (,),
	// for example, domainName1.userName1,do mainName2.userName2.
	PrincipalName string `json:"principal_name"`
	// Authorization operation type.
	// - putRecords: upload data.
	// - getRecords: download data.
	// Enumeration values:
	// putRecords
	// getRecords
	ActionType string `json:"action_type"`
	// Authorization impact type.
	// - accept: The authorization operation is allowed.
	// Enumeration values:
	// - accept
	Effect string `json:"effect"`
}

func CreatePolicyRule(client *golangsdk.ServiceClient, opts CreatePolicyRuleOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v2/{project_id}/streams/{stream_name}/policies
	_, err = client.Post(client.ServiceURL("streams", opts.StreamName, "policies"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
