package loggroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, groupId string) (*GetResponse, error) {
	// GET /v2.0/{project_id}/log-groups/{group_id}
	raw, err := client.Get(client.ServiceURL("log-groups", groupId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetResponse struct {
	// Log group ID
	LogGroupId string `json:"log_group_id"`
	// Log group name
	LogGroupName string `json:"log_group_name"`
	// Log group creation time
	CreationTime int64 `json:"creation_time"`
	// Log expiration time
	TTLInDays int `json:"ttl_in_days"`
}
