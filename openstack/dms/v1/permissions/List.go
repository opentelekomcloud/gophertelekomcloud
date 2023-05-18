package permissions

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient, instanceId, topicName string) (*Permissions, error) {
	// GET /v2/{project_id}/instances/{instance_id}/users

	url := client.ServiceURL("instances", instanceId, "topics", topicName, "accesspolicy")
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res Permissions
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Permissions struct {
	Name      string   `json:"name"`
	TopicType int      `json:"topic_type"`
	Policies  []Policy `json:"policies"`
}

type Policy struct {
	Owner        bool   `json:"owner"`
	UserName     string `json:"user_name"`
	AccessPolicy string `json:"access_policy"`
}
