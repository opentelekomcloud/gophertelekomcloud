package connection

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type KillSessionsOpts struct {
	Sessions []string `json:"sessions" required:"true"`
}

func KillSessions(client *golangsdk.ServiceClient, nodeId string, opts KillSessionsOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}
	_, err = client.Post(client.ServiceURL("nodes", nodeId, "session"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return err
	}
	return nil
}
