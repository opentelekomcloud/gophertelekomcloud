package messages

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ShowInstanceMessagesOpts struct {
	Topic         string `q:"topic,required"`
	Asc           *bool  `q:"asc,omitempty"`
	StartTime     string `q:"start_time,omitempty"`
	EndTime       string `q:"end_time,omitempty"`
	Limit         string `q:"limit,omitempty"`
	Offset        string `q:"offset,omitempty"`
	Download      *bool  `q:"download,omitempty"`
	MessageOffset string `q:"message_offset,omitempty"`
	Partition     string `q:"partition,omitempty"`
	Keyword       string `q:"keyword,omitempty"`
}

// ShowInstanceMessages sends GET /v2/{project_id}/instances/{instance_id}/messages
func ShowInstanceMessages(client *golangsdk.ServiceClient, instanceID string, opts ShowInstanceMessagesOpts) (*ShowInstanceMessagesResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "messages").
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

	var res ShowInstanceMessagesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ShowInstanceMessagesResponse struct {
	Messages []MessagesEntity `json:"messages"`
	Total    int64            `json:"total"`
	Size     int64            `json:"size"`
}

type MessagesEntity struct {
	Topic         string `json:"topic"`
	Partition     int    `json:"partition"`
	Key           string `json:"key"`
	Value         string `json:"value"`
	Size          int    `json:"size"`
	Timestamp     int64  `json:"timestamp"`
	HugeMessage   bool   `json:"huge_message"`
	MessageOffset int64  `json:"message_offset"`
	MessageID     string `json:"message_id"`
	AppID         string `json:"app_id"`
	Tag           string `json:"tag"`
}
