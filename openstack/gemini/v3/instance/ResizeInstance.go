package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ResizeInstanceOpts struct {
	InstanceID     string `json:"-"`
	TargetSpecCode string `json:"-" required:"true"`
}

func ResizeInstance(client *golangsdk.ServiceClient, opts ResizeInstanceOpts) (string, error) {
	body := map[string]interface{}{
		"resize": map[string]interface{}{
			"target_spec_code": opts.TargetSpecCode,
		},
	}

	b, err := build.RequestBody(body, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Put(client.ServiceURL("instances", opts.InstanceID, "resize"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		JobId string `json:"job_id"`
	}

	err = extract.Into(raw.Body, &res)
	if err != nil {
		return "", err
	}

	return res.JobId, nil
}
