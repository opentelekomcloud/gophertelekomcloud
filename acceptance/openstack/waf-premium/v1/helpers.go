package v1

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf-premium/v1/instances"
)

func waitForInstanceToBeCreated(client *golangsdk.ServiceClient, secs int, id string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		instance, err := instances.Get(client, id)
		if err != nil {
			return false, err
		}
		if instance.Status == 1 {
			return true, nil
		}
		if instance.Status == 4 {
			return false, fmt.Errorf("error creating instance")
		}

		return false, nil
	})
}
