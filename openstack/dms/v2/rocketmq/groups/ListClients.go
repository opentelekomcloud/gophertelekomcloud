package groups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListClientsOpts struct {
	// Number of records to query. Range: 1-50.
	Limit int `q:"limit,omitempty"`
	// Offset where the query starts. Must be greater than or equal to 0.
	Offset int `q:"offset,omitempty"`
	// Whether to query the detailed consumer list.
	IsDetail bool `q:"is_detail,omitempty"`
}

// ListClientsResponse is returned by ListClients.
type ListClientsResponse struct {
	// Consumer group name.
	GroupName string `json:"group_name"`
	// Whether the consumer group is online.
	Online bool `json:"online"`
	// Whether subscriptions are consistent.
	SubscriptionConsistency bool `json:"subscription_consistency"`
	// Total number of consumers.
	Total int `json:"total"`
	// Offset of the next page.
	NextOffset int `json:"next_offset"`
	// Offset of the previous page.
	PreviousOffset int `json:"previous_offset"`
	// Consumer subscription detail list.
	Clients []Client `json:"clients"`
}

type Client struct {
	// Client language.
	Language string `json:"language"`
	// Client version.
	Version string `json:"version"`
	// Client ID.
	ClientID string `json:"client_id"`
	// Client address.
	ClientAddr string `json:"client_addr"`
	// Subscription list.
	Subscriptions []Subscription `json:"subscriptions"`
}

type Subscription struct {
	// Name of the subscribed topic.
	Topic string `json:"topic"`
	// Subscription type.
	//    TAG: tag-based subscription.
	//    SQL92: message-attribute-based subscription.
	Type string `json:"type"`
	// Subscription tag.
	Expression string `json:"expression"`
}

// ListClients queries the consumer list of a consumer group.
// Send GET to /v2/{project_id}/rocketmq/instances/{instance_id}/groups/{group}/clients
func ListClients(client *golangsdk.ServiceClient, instanceID, group string, opts ListClientsOpts) (*ListClientsResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("rocketmq", "instances", instanceID, "groups", group, "clients").
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

	var res ListClientsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
