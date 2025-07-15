package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Delete given topics belong to the instance id
func Delete(client *golangsdk.ServiceClient, instanceID string, topics []string) (*DeleteResponse, error) {
	var delOpts = struct {
		Topics []string `json:"topics" required:"true"`
	}{Topics: topics}

	b, err := build.RequestBody(delOpts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", instanceID, "topics", "delete"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res DeleteResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// DeleteResponse is a struct that contains the deletion response
type DeleteResponse struct {
	Topics []TopicDelete `json:"topics"`
}

type TopicDelete struct {
	Name    string `json:"id"`
	Success bool   `json:"success"`
}
