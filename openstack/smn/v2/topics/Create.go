package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts is a struct that contains all the parameters.
type CreateOpts struct {
	// Name of the topic to be created
	Name string `json:"name" required:"true"`

	// Topic display name
	DisplayName string `json:"display_name,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CreateResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/notifications/topics
	raw, err := client.Post(client.ServiceURL("topics"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201, 200},
	})
	if err != nil {
		return nil, err
	}

	var res CreateResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CreateResp struct {
	RequestId string `json:"request_id"`
	TopicUrn  string `json:"topic_urn"`
}
