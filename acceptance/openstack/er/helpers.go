package er

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/er/v3/instance"
)

func waitForInstanceAvailable(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		erInstance, err := instance.Get(client, instanceID)
		if err != nil {
			return false, err
		}
		if erInstance.Instance.State == "available" {
			return true, nil
		}
		return false, nil
	})
}

func waitForInstanceDeleted(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := instance.Get(client, instanceID)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		return false, nil
	})
}
