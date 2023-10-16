package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func Start(client *golangsdk.ServiceClient, id string) (*JobId, error) {
	type Start struct {
		Start *EmptyObj `json:"start"`
	}

	// POST /v1.1/{project_id}/clusters/{cluster_id}/action
	raw, err := client.Post(client.ServiceURL("clusters", id, "action"), Start{}, nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en"},
	})
	return extraJob(err, raw)
}

type EmptyObj struct {
	Obj *string `json:"-"`
}
