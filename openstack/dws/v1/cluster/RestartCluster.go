package cluster

import (
	"fmt"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type RestartClusterOpts struct {
	ClusterId string `json:"-"`
	// Restart flag.
	Restart interface{} `json:"restart" required:"true"`
}

func RestartCluster(client *golangsdk.ServiceClient, opts RestartClusterOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// POST /v1.0/{project_id}/clusters/{cluster_id}/restart
	_, err = client.Post(client.ServiceURL("clusters", opts.ClusterId, "restart"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func WaitForRestart(c *golangsdk.ServiceClient, id string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		current, err := ListClusterDetails(c, id)
		if err != nil {
			return false, err
		}

		if current.Status == "AVAILABLE" && current.TaskStatus != "REBOOTING" {
			return true, nil
		}

		if current.TaskStatus == "REBOOT_FAILURE" {
			return false, fmt.Errorf("cluster Restart failed: " + current.FailedReasons.ErrorMsg)
		}

		time.Sleep(10 * time.Second)

		return false, nil
	})
}
