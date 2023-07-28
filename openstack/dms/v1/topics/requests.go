package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateOptsBuilder interface {
	ToTopicCreateMap() (map[string]interface{}, error)
}

type DeleteOptsBuilder interface {
	ToTopicDeleteMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	Name             string `json:"id" required:"true"`
	Partition        int    `json:"partition,omitempty"`
	Replication      int    `json:"replication,omitempty"`
	SyncReplication  bool   `json:"sync_replication,omitempty"`
	RetentionTime    int    `json:"retention_time,omitempty"`
	SyncMessageFlush bool   `json:"sync_message_flush,omitempty"`
}

func (opts CreateOpts) ToTopicCreateMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, "")
}

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder, instanceId string) (r CreateResult) {
	b, err := opts.ToTopicCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client, instanceId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func Get(client *golangsdk.ServiceClient, instanceId string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, instanceId), &r.Body, nil)
	return
}

type DeleteOpts struct {
	Topics []string `json:"topics" required:"true"`
}

func (opts DeleteOpts) ToTopicDeleteMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, "")
}

func Delete(client *golangsdk.ServiceClient, opts DeleteOptsBuilder, instanceId string) (r DeleteResult) {
	b, err := opts.ToTopicDeleteMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(deleteURL(client, instanceId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
