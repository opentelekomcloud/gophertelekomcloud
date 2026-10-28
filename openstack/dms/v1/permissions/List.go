package permissions

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List sends GET /v1/{project_id}/instances/{instance_id}/topics/{topic_name}/accesspolicy
func List(client *golangsdk.ServiceClient, instanceId, topicName string) (*Permissions, error) {
	if instanceId == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "instanceId"}
	}
	if topicName == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "topicName"}
	}

	raw, err := client.Get(client.ServiceURL("instances", instanceId, "topics", topicName, "accesspolicy"), nil, nil)
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
