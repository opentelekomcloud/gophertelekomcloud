package migration

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetOpts struct {
	// Query type: vhost, exchange or queue.
	// Mandatory for rabbitToRocket migration tasks, ignored for rocketmq ones.
	Type string `q:"type,omitempty"`
	// Virtual host name.
	//    Can be left empty when the vhost list is queried.
	//    Name of the vhost the exchange belongs to when the exchange list is queried.
	//    {vhost}-{exchange} of the queue when the queue list is queried.
	Name string `q:"name,omitempty"`
	// Current page, starting from 1.
	Offset int `q:"offset,omitempty"`
	// Current page size.
	Limit int `q:"limit,omitempty"`
}

type getQuery struct {
	ID     string `q:"id"`
	Type   string `q:"type,omitempty"`
	Name   string `q:"name,omitempty"`
	Offset int    `q:"offset,omitempty"`
	Limit  int    `q:"limit,omitempty"`
}

// TaskDetails contains the details of a metadata migration task.
type TaskDetails struct {
	Task
	// Metadata submitted with the task, as a JSON string.
	JSONContent string `json:"json_content"`
	// Migration result, as a JSON string.
	Result string `json:"result"`
}

// Get the details of a specified metadata migration task of a RocketMQ instance.
// Send GET to /v2/{project_id}/instances/{instance_id}/metadata?id={task_id}
func Get(client *golangsdk.ServiceClient, instanceID, taskID string, opts GetOpts) (*TaskDetails, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "metadata").
		WithQueryParams(&getQuery{
			ID:     taskID,
			Type:   opts.Type,
			Name:   opts.Name,
			Offset: opts.Offset,
			Limit:  opts.Limit,
		}).
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

	var res TaskDetails
	err = extract.Into(raw.Body, &res)
	return &res, err
}
