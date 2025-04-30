package log

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(c *golangsdk.ServiceClient, id string) (err error) {
	// DELETE /v3/{project_id}/elb/logtanks/{logtank_id}
	_, err = c.Delete(c.ServiceURL("logtanks", id), nil)
	return
}
