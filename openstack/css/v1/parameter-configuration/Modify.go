package parameter_configuration

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ModifyOpts struct {
	// Operations performed on parameter configurations. The value can be:
	// modify
	// delete
	// reset
	Edit map[string]interface{} `json:"edit" required:"true"`
}

func Modify(client *golangsdk.ServiceClient, opts ModifyOpts, clusterID string) (*Config, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	// POST /v1.0/{project_id}/clusters/{cluster_id}/ymls/update
	raw, err := client.Post(client.ServiceURL("clusters", clusterID, "ymls", "update"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Config
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Config struct {
	// Indicates whether the modification is successful.
	// true: The modification is successful.
	// false: The modification failed.
	Acknowledged bool `json:"acknowledged"`
	// Error message. If acknowledged is set to true, null is returned for this field.
	ExternalMessage string `json:"externalMessage"`
	// HTTP error information. The default value is null.
	ErrorMsg string `json:"httpErrorResponse"`
}
