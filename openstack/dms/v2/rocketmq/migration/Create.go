package migration

import (
	"strconv"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Whether to overwrite the configurations of existing topics and groups
	// with the same name. If false, an error is reported when a topic or
	// group already exists.
	Overwrite bool `json:"-"`
	// Migration task name.
	Name string `json:"-" required:"true"`
	// Migration task type: rocketmq or rabbitToRocket.
	Type string `json:"-" required:"true"`
	// Topic metadata keyed by topic name.
	// Mandatory for rocketmq migration tasks.
	TopicConfigTable map[string]TopicConfig `json:"topic_config_table,omitempty"`
	// Consumer group metadata keyed by consumer group name.
	// Mandatory for rocketmq migration tasks.
	SubscriptionGroupTable map[string]SubscriptionGroup `json:"subscription_group_table,omitempty"`
	// RabbitMQ virtual host metadata.
	// Mandatory for rabbitToRocket migration tasks.
	Vhosts []Vhost `json:"vhosts,omitempty"`
	// RabbitMQ queue metadata.
	// Mandatory for rabbitToRocket migration tasks.
	Queues []Queue `json:"queues,omitempty"`
	// RabbitMQ exchange metadata.
	// Mandatory for rabbitToRocket migration tasks.
	Exchanges []Exchange `json:"exchanges,omitempty"`
	// RabbitMQ binding metadata.
	// Mandatory for rabbitToRocket migration tasks.
	Bindings []Binding `json:"bindings,omitempty"`
}

type createQuery struct {
	Overwrite string `q:"overwrite"`
	Name      string `q:"name"`
	Type      string `q:"type"`
}

// CreateResponse is returned by Create.
type CreateResponse struct {
	// Task ID.
	TaskID string `json:"task_id"`
}

// Create a metadata migration task of a RocketMQ instance.
// Send POST to /v2/{project_id}/instances/{instance_id}/metadata
func Create(client *golangsdk.ServiceClient, instanceID string, opts CreateOpts) (*CreateResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "metadata").
		WithQueryParams(&createQuery{
			Overwrite: strconv.FormatBool(opts.Overwrite),
			Name:      opts.Name,
			Type:      opts.Type,
		}).
		Build()
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
