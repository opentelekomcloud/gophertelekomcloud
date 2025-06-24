package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get a topic with detailed information by id
func Get(client *golangsdk.ServiceClient, id string) (*Topic, error) {
	raw, err := client.Get(client.ServiceURL("topics", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Topic
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Topic struct {
	TopicUrn    string `json:"topic_urn"`
	DisplayName string `json:"display_name"`
	Name        string `json:"name"`
	PushPolicy  int    `json:"push_policy"`
	UpdateTime  string `json:"update_time"`
	CreateTime  string `json:"create_time"`
}
