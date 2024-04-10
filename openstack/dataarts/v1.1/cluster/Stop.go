package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

const (
	StopImmediately = "IMMEDIATELY"
	StopGracefully  = "GRACEFULLY"
)

type StopOpts struct {
	// Stop Cluster stop operation, which defines the parameters for stopping a cluster.
	Stop StopStruct `json:"stop" required:"true"`
}

type StopStruct struct {
	// StopMode should be StopImmediately or StopGracefully
	StopMode string `json:"stopMode,omitempty"`
	// DelayTime Stop delay, in seconds.
	// This parameter is valid only when stopMode is set to GRACEFULLY.
	// If the value of this parameter is set to -1, the system waits for all jobs to complete and stops accepting new jobs.
	// If the value of this parameter is greater than 0, the system stops the cluster after the specified time and stops accepting new jobs.
	DelayTime int `json:"delayTime,omitempty"`
}

// Stop is used to stop a cluster.
// Send request POST /v1.1/{project_id}/clusters/{cluster_id}/action
func Stop(client *golangsdk.ServiceClient, clusterId string, opts StopOpts) (*JobId, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(
		client.ServiceURL("clusters", clusterId, "action"),
		b,
		nil,
		&golangsdk.RequestOpts{
			MoreHeaders: map[string]string{"Content-Type": "application/json", HeaderXLanguage: RequestedLang},
		})
	return respToJobId(raw)
}
