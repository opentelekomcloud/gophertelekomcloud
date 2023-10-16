package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type RestartOpts struct {
	Id      string        `json:"-"`
	Restart RestartStruct `json:"restart" required:"true"`
}

type RestartStruct struct {
	StopMode     int    `json:"restartDelayTime,omitempty"`
	RestartMode  string `json:"restartMode,omitempty"`
	RestartLevel string `json:"restartLevel,omitempty"`
	Type         string `json:"type,omitempty"`
	Instance     string `json:"instance,omitempty"`
	Group        string `json:"group,omitempty"`
}

func Restart(client *golangsdk.ServiceClient, opts RestartOpts) (*JobId, error) {
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
