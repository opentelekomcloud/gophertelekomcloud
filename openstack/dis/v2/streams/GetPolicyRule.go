package streams

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetPolicyRule(client *golangsdk.ServiceClient, streamName string) (*GetPolicyRuleResponse, error) {
	// GET /v2/{project_id}/streams/{stream_name}/policies
	raw, err := client.Get(client.ServiceURL("streams", streamName, "policies"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetPolicyRuleResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetPolicyRuleResponse struct {
	// Unique ID of the stream.
	StreamId string `json:"stream_id,omitempty"`
	// List of authorization information records.
	Rules []PrincipalRule `json:"rules,omitempty"`
}

type PrincipalRule struct {
	// ID of the authorized user.
	Principal string `json:"principal,omitempty"`
	// Name of the authorized user.
	// If the permission is granted to all sub-users of a tenant, the format is domainName.*.
	// If the permission is granted to a specified sub-user of a tenant, the format is domainName.userName.
	PrincipalName string `json:"principal_name,omitempty"`
	// Authorization operation type.
	// - putRecords: the data to be uploaded.
	// - getRecords: Download data.
	// Enumeration values:
	// putRecords
	// getRecords
	ActionType string `json:"action_type,omitempty"`
	// Authorization impact type.
	// - accept: The authorization operation is allowed.
	// Enumeration values:
	// accept
	Effect string `json:"effect,omitempty"`
}
