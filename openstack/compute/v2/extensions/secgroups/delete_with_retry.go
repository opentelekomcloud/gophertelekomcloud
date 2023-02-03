package secgroups

import (
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// DeleteWithRetry will try to permanently delete a particular security
// group based on its unique ID and RetryTimeout.
func DeleteWithRetry(client *golangsdk.ServiceClient, id string, timeout int) error {
	return golangsdk.WaitFor(timeout, func() (bool, error) {
		_, err := client.Delete(client.ServiceURL("os-security-groups", id), nil)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault400); ok {
				time.Sleep(10 * time.Second)
				return false, nil
			}
			return false, err
		}
		return true, nil
	})
}
