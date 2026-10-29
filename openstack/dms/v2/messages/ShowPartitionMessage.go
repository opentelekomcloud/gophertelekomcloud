package messages

import (
	"strconv"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ShowPartitionMessageOpts struct {
	MessageOffset string `q:"message_offset,required"`
}

// ShowPartitionMessage sends GET /v2/{project_id}/instances/{instance_id}/management/topics/{topic}/partitions/{partition}/message
func ShowPartitionMessage(client *golangsdk.ServiceClient, instanceID, topic string, partition int, opts ShowPartitionMessageOpts) (*ShowPartitionMessageResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "management", "topics", topic, "partitions", strconv.Itoa(partition), "message").
		WithQueryParams(&opts).
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

	var res ShowPartitionMessageResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ShowPartitionMessageResponse struct {
	Message []ShowPartitionMessageEntity `json:"message"`
}

type ShowPartitionMessageEntity struct {
	Key           string `json:"key"`
	Value         string `json:"value"`
	Topic         string `json:"topic"`
	Partition     int    `json:"partition"`
	MessageOffset int64  `json:"message_offset"`
	Size          int    `json:"size"`
	Timestamp     int64  `json:"timestamp"`
}
