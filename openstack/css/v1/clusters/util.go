package clusters

import (
	"fmt"
	"log"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func WaitForClusterOperationSucces(client *golangsdk.ServiceClient, id string, timeout int) error {
	return golangsdk.WaitFor(timeout, func() (bool, error) {
		cluster, err := Get(client, id)
		if err != nil {
			if _, ok := err.(golangsdk.BaseError); ok {
				return true, err
			}
			log.Printf("Error waiting for CSS cluster: %s", err) // ignore connection-related errors
			return false, nil
		}

		switch s := cluster.Status; s {
		case "100":
			time.Sleep(30 * time.Second) // make a bigger wait if it's not ready
			return false, nil
		case "200":
			return true, nil
		case "303":
			return true, fmt.Errorf("cluster operartion failed: %+v", cluster.FailedReasons)
		default:
			return true, fmt.Errorf("invalid status: %s", s)
		}
	})
}

func WaitForClusterToExtend(client *golangsdk.ServiceClient, id string, timeout int) error {
	return golangsdk.WaitFor(timeout, func() (bool, error) {
		cluster, err := Get(client, id)
		if err != nil {
			if _, ok := err.(golangsdk.BaseError); ok {
				return true, err
			}
			log.Printf("Error waiting for CSS cluster to extend: %s", err) // ignore connection-related errors
			return false, nil
		}
		// No active action
		if len(cluster.Actions) == 0 {
			return true, nil
		}
		if cluster.Actions[0] == "GROWING" || cluster.Actions[0] == "RESIZING_VOLUME" {
			time.Sleep(30 * time.Second) // make a bigger wait if it's not ready
			return false, nil
		}
		return false, fmt.Errorf("unexpected cluster actions: %v; progress: %v", cluster.Actions, cluster.ActionProgress)
	})
}
