package messages

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ShowMessagesOpts struct {
	StartTime string `q:"start_time,omitempty"`
	EndTime   string `q:"end_time,omitempty"`
	Limit     int    `q:"limit,omitempty"`
	Offset    *int   `q:"offset,omitempty"`
	Partition string `q:"partition,omitempty"`
}

// ShowMessages sends GET /v2/{project_id}/instances/{instance_id}/management/topics/{topic}/messages
func ShowMessages(client *golangsdk.ServiceClient, instanceID, topic string, opts ShowMessagesOpts) (*ShowMessagesResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "management", "topics", topic, "messages").
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

	var res ShowMessagesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ShowMessagesResponse struct {
	Messages      []Messages `json:"messages"`
	MessagesCount int        `json:"messages_count"`
	OffsetsCount  int        `json:"offsets_count"`
	Offset        int        `json:"offset"`
}

type Messages struct {
	Topic         string `json:"topic"`
	Partition     int    `json:"partition"`
	MessageOffset int64  `json:"message_offset"`
	Size          int    `json:"size"`
	Timestamp     int64  `json:"timestamp"`
}
