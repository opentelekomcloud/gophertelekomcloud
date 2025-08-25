package backup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	InstanceId  string `json:"instance_id" required:"true"`
	Name        string `json:"name" required:"true"`
	Description string `json:"description,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CreateResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("backups", "create"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}

	var res CreateResponse
	return &res, extract.Into(raw.Body, &res)
}

type CreateResponse struct {
	Backup BackupInfo `json:"backup"`
	JobId  string     `json:"job_id"`
}

type BackupInfo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	BeginTime   string `json:"begin_time"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	InstanceId  string `json:"instance_id"`
}
