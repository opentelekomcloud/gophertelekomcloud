package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type StartOpts struct {
	// Start Starting a cluster. This parameter is an empty object.
	Start *EmptyStruct `json:"start"`
}

type EmptyStruct struct{}

// Start is used to start a cluster.
// Send request POST /v1.1/{project_id}/clusters/{cluster_id}/action
func Start(client *golangsdk.ServiceClient, clusterId string, startOpts *StartOpts) (*JobId, error) {
	b, err := build.RequestBody(startOpts, "")
	if err != nil {
		return nil, err
	}

	resp, err := client.Post(
		client.ServiceURL("clusters", clusterId, "action"),
		b,
		nil,
		&golangsdk.RequestOpts{
			MoreHeaders: map[string]string{"Content-Type": "application/json", HeaderXLanguage: RequestedLang},
		},
	)
	if err != nil {
		return nil, err
	}

	return respToJobId(resp)
}
