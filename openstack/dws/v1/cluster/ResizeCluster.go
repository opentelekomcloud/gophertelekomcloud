package cluster

import (
	"fmt"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ResizeClusterOpts struct {
	ClusterId string `json:"-"`
	// Number of nodes to be added
	Count int `json:"count" required:"true"`
}

func ResizeCluster(client *golangsdk.ServiceClient, opts ResizeClusterOpts) (err error) {
	b, err := build.RequestBody(opts, "scale_out")
	if err != nil {
		return
	}

	// POST /v1.0/{project_id}/clusters/{cluster_id}/resize
	_, err = client.Post(client.ServiceURL("clusters", opts.ClusterId, "resize"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func WaitForResize(c *golangsdk.ServiceClient, id string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		current, err := ListClusterDetails(c, id)
		if err != nil {
			return false, err
		}

		if current.Status == "AVAILABLE" && current.TaskStatus != "GROWING" {
			return true, nil
		}

		if current.TaskStatus == "RESIZE_FAILURE" {
			return false, fmt.Errorf("cluster RESIZE failed: " + current.FailedReasons.ErrorMsg)
		}

		time.Sleep(10 * time.Second)

		return false, nil
	})
}
