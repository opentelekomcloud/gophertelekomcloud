package messages

import (
	"strconv"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ShowPartitionBeginningMessage sends GET /v2/{project_id}/instances/{instance_id}/management/topics/{topic}/partitions/{partition}/beginning-message
func ShowPartitionBeginningMessage(client *golangsdk.ServiceClient, instanceID, topic string, partition int) (*ShowPartitionBeginningMessageResponse, error) {
	if instanceID == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "instanceID"}
	}
	if topic == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "topic"}
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "management", "topics", topic, "partitions", strconv.Itoa(partition), "beginning-message").
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ShowPartitionBeginningMessageResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ShowPartitionBeginningMessageResponse struct {
	Topic     string `json:"topic"`
	Partition int    `json:"partition"`
	Offset    int    `json:"offset"`
	Timestamp int64  `json:"timestamp"`
}
