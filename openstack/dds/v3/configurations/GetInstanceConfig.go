package configurations

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ConfigOpts struct {
	// Instance ID, group ID, or node ID. You can call the API used for querying instances and details to obtain the value.
	// If you do not have an instance, you can call the API used for creating an instance.
	// If the DB instance type is cluster and the shard or config parameter template is obtained, the value is the group ID.
	// If the parameter template of the mongos node is obtained, the value is the node ID.
	// If the DB instance type is a replica set instance or a single node instance, the value is the instance ID.
	EntityId string `q:"entity_id"`
}

func GetInstanceConfig(client *golangsdk.ServiceClient, instanceId string, opts ConfigOpts) (*Response, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceId, "configurations").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3/{project_id}/instances/{instance_id}/configurations
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res Response
	err = extract.Into(raw.Body, &res)
	return &res, err
}
