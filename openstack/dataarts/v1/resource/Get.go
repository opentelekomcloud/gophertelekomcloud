package resource

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get is used to query resource details. A resource contains various files such as JAR, ZIP, and properties files. A created resource can be used in job nodes such as DLI Spark and MRS Spark.
// Send request GET /v1/{project_id}/resources/{resource_id}
func Get(client *golangsdk.ServiceClient, resourceId, workspace string) (*Resource, error) {

	var opts golangsdk.RequestOpts
	if workspace != "" {
		opts.MoreHeaders = map[string]string{HeaderWorkspace: workspace}
	}

	raw, err := client.Get(client.ServiceURL(resourcesEndpoint, resourceId), nil, &opts)
	if err != nil {
		return nil, err
	}

	var res Resource
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
