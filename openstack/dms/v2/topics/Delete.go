package topics

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Delete given topics belong to the instance id.
// If only some topics are deleted, a success response is returned; check Topics[].Success for failed topics.
// Send POST /v2/{project_id}/instances/{instance_id}/topics/delete
func Delete(client *golangsdk.ServiceClient, instanceID string, topics []string) (*DeleteResponse, error) {
	if instanceID == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "instanceID"}
	}

	// Topics to delete. The API reference marks it as optional, but it is mandatory for batch deletion.
	// The required tag only rejects nil, so empty lists and empty names are checked explicitly.
	if len(topics) == 0 {
		return nil, golangsdk.ErrMissingInput{Argument: "topics"}
	}
	for i, name := range topics {
		if name == "" {
			return nil, golangsdk.ErrMissingInput{Argument: fmt.Sprintf("topics[%d]", i)}
		}
	}

	var delOpts = struct {
		Topics []string `json:"topics"`
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
	// Topic list.
	Topics []TopicDelete `json:"topics"`
}

// TopicDelete represents the deletion result of one topic
type TopicDelete struct {
	// Topic name.
	Name string `json:"id"`
	// Whether the topics are deleted.
	Success bool `json:"success"`
}
