package topics

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOps is a struct that contains all the parameters of create function
type CreateOpts struct {
	// the name/ID of a topic
	Name string `json:"id" required:"true"`
	// topic partitions, value range: 1-100
	Partition int `json:"partition,omitempty"`
	// topic replications, value range: 1-3
	Replication int `json:"replication,omitempty"`
	// aging time in hours, value range: 1-168, defaults to 72
	RetentionTime int `json:"retention_time,omitempty"`

	SyncMessageFlush bool `json:"sync_message_flush,omitempty"`
	SyncReplication  bool `json:"sync_replication,omitempty"`
}

// Create a kafka topic with given parameters
func Create(client *golangsdk.ServiceClient, instanceID string, opts CreateOpts) (*CreateResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(rootURL(client, instanceID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res CreateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// UpdateOpts is a struct which represents the parameters of update function
type UpdateOpts struct {
	Topics []UpdateItem `json:"topics" required:"true"`
}

// UpdateItem represents the object of one topic in update function
type UpdateItem struct {
	// Name can not be updated
	Name             string `json:"id" required:"true"`
	Partition        *int   `json:"new_partition_numbers,omitempty"`
	RetentionTime    *int   `json:"retention_time,omitempty"`
	SyncMessageFlush *bool  `json:"sync_message_flush,omitempty"`
	SyncReplication  *bool  `json:"sync_replication,omitempty"`
}

// Update is a method which can be able to update topics
func Update(client *golangsdk.ServiceClient, instanceID string, opts UpdateOpts) error {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}
	_, err = client.Put(rootURL(client, instanceID), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}

// Get an topic with detailed information by instance id and topic name
func Get(client *golangsdk.ServiceClient, instanceID, topicName string) (*TopicDetail, error) {
	raw, err := client.Get(getURL(client, instanceID, topicName), nil, nil)
	if err != nil {
		return nil, err
	}

	var res TopicDetail
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// List all topics belong to the instance id
func List(client *golangsdk.ServiceClient, instanceID string) (*ListResponse, error) {
	raw, err := client.Get(rootURL(client, instanceID), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// Delete given topics belong to the instance id
func Delete(client *golangsdk.ServiceClient, instanceID string, topics []string) (*DeleteResponse, error) {
	var delOpts = struct {
		Topics []string `json:"topics" required:"true"`
	}{Topics: topics}

	b, err := build.RequestBody(delOpts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(deleteURL(client, instanceID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res DeleteResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
