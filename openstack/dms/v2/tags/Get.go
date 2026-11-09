package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get sends GET /v2/{project_id}/kafka/{instance_id}/tags
func Get(client *golangsdk.ServiceClient, instanceID string) (*GetResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("kafka", instanceID, "tags").
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetResponse struct {
	Tags []Tag `json:"tags"`
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
