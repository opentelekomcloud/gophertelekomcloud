package addons

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// WaitForAddonRunning - wait until addon status is `running`
func WaitForAddonRunning(client *golangsdk.ServiceClient, id, clusterID string, timeoutSeconds int) error {
	return golangsdk.WaitFor(timeoutSeconds, func() (bool, error) {
		addon, err := Get(client, id, clusterID)
		if err != nil {
			return false, fmt.Errorf("error retriving addon status: %w", err)
		}
		if addon.Status.Status == "running" {
			return true, nil
		}
		return false, nil
	})
}

// WaitForAddonDeleted - wait until addon is deleted
func WaitForAddonDeleted(client *golangsdk.ServiceClient, id, clusterID string, timeoutSeconds int) error {
	return golangsdk.WaitFor(timeoutSeconds, func() (bool, error) {
		_, err := Get(client, id, clusterID)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, fmt.Errorf("error retriving addon status: %w", err)
		}
		return false, nil
	})
}
