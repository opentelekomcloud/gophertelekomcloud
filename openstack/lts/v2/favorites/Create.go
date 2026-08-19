package favorites

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// EnterpriseProjectID is the enterprise project ID.
	EnterpriseProjectID string `json:"eps_id,omitempty"`
	// ResourceID is the favorite resource ID.
	ResourceID string `json:"favorite_resource_id" required:"true"`
	// ResourceType is the favorite resource type, LOG_STREAM or LOG_GROUP.
	ResourceType string `json:"favorite_resource_type" required:"true"`
	// LogGroupID is the log group ID.
	LogGroupID string `json:"log_group_id" required:"true"`
	// LogGroupName is the log group name.
	LogGroupName string `json:"log_group_name,omitempty"`
	// LogStreamID is the log stream ID.
	LogStreamID string `json:"log_stream_id" required:"true"`
	// LogStreamName is the log stream name.
	LogStreamName string `json:"log_stream_name,omitempty"`
	// IsGlobal indicates whether global favorites are enabled. The API requires true.
	IsGlobal bool `json:"is_global" required:"true"`
}

type Favorite struct {
	CreateTime          int64  `json:"create_time"`
	EnterpriseProjectID string `json:"eps_id"`
	ResourceID          string `json:"favorite_resource_id"`
	ResourceType        string `json:"favorite_resource_type"`
	LogGroupID          string `json:"log_group_id"`
	LogGroupName        string `json:"log_group_name"`
	LogStreamID         string `json:"log_stream_id"`
	LogStreamName       string `json:"log_stream_name"`
	ProjectID           string `json:"project_id"`
	IsGlobal            bool   `json:"is_global"`
	LogGroupNameAlias   string `json:"log_group_name_alias"`
	LogStreamNameAlias  string `json:"log_stream_name_alias"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Favorite, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1.0/{project_id}/lts/favorite
	raw, err := client.Post(client.ServiceURL("lts", "favorite"), b, nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"Content-Type": "application/json;charset=utf8",
		},
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res Favorite
	err = extract.Into(raw.Body, &res)
	return &res, err
}
