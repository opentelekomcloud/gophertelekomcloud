package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListProducersOpts represents the query parameters of the producer list request
type ListProducersOpts struct {
	// Offset. The records after this offset will be queried.
	Offset *int `q:"offset,omitempty"`
	// Maximum number of records that can be returned.
	Limit *int `q:"limit,omitempty"`
}

// ListProducers queries the current producer list of a topic
// Send GET /v2/{project_id}/kafka/instances/{instance_id}/topics/{topic}/producers
func ListProducers(client *golangsdk.ServiceClient, instanceID, topicName string, opts ListProducersOpts) (*ListProducersResponse, error) {
	if instanceID == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "instanceID"}
	}
	if topicName == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "topicName"}
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("kafka", "instances", instanceID, "topics", topicName, "producers").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListProducersResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// ListProducersResponse is a struct that contains the producer list response
type ListProducersResponse struct {
	// Total records.
	Total int `json:"total"`
	// Producer list.
	Producers []Producer `json:"producers"`
}

// Producer represents one producer of a topic
type Producer struct {
	// Producer address.
	ProducerAddress string `json:"producer_address"`
	// Broker address.
	BrokerAddress string `json:"broker_address"`
	// Time when the broker was connected.
	JoinTime int64 `json:"join_time"`
}
