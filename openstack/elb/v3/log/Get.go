package log

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Log, error) {
	// GET /v3/{project_id}/elb/logtanks/{logtank_id}
	raw, err := client.Get(client.ServiceURL("logtanks", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Log
	err = extract.Into(raw.Body, &res)
	return &res, err
}
