package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Topic name. Enter 3 to 64 characters. Use only letters, digits,
	// percent (%), vertical bars (|), hyphens (-), and underscores (_).
	Name string `json:"name" required:"true"`
	// Total number of queues. Range: 1-50.
	QueueNum int `json:"queue_num,omitempty"`
	// Message type. Mandatory only for RocketMQ 5.x instances.
	//    NORMAL: normal messages.
	//    FIFO: ordered messages.
	//    DELAY: scheduled messages.
	//    TRANSACTION: transactional messages.
	MessageType string `json:"message_type,omitempty"`
}

// CreateResponse is returned by Create.
type CreateResponse struct {
	// Topic name.
	ID string `json:"id"`
}

// Create a topic of a RocketMQ instance.
// Send POST to /v2/{project_id}/instances/{instance_id}/topics
func Create(client *golangsdk.ServiceClient, instanceID string, opts CreateOpts) (*CreateResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceID, "topics").Build()
	if err != nil {
		return nil, err
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res CreateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
