package cluster

import (
	"net/http"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type StopOpts struct {
	Id   string     `json:"-"`
	Stop StopStruct `json:"stop" required:"true"`
}

type StopStruct struct {
	StopMode  string `json:"stopMode,omitempty"`
	DelayTime string `json:"delayTime,omitempty"`
}

func Stop(client *golangsdk.ServiceClient, opts StopOpts) (*JobId, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	// POST /v1.1/{project_id}/clusters/{cluster_id}/action
	raw, err := client.Post(client.ServiceURL("clusters", opts.Id, "action"), b, nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en"},
	})
	return extraJob(err, raw)
}

type JobId struct {
	JobId []string `json:"jobId"`
}

func extraJob(err error, raw *http.Response) (*JobId, error) {
	if err != nil {
		return nil, err
	}

	var res JobId
	err = extract.Into(raw.Body, &res)
	return &res, err
}
