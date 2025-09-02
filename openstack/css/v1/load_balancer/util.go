package load_balancer

import (
	"fmt"
	"log"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func WaitForListenerStatus(client *golangsdk.ServiceClient, id string, timeout int) error {
	time.Sleep(5 * time.Second)
	return golangsdk.WaitFor(timeout, func() (bool, error) {
		elb, err := Get(client, id)
		if err != nil {
			if _, ok := err.(golangsdk.BaseError); ok {
				return true, err
			}
			log.Printf("Error while waiting for listener's status to change to online: %s", err)
			return false, nil
		}
		if len(elb.Healthmonitors) == 0 {
			return false, fmt.Errorf("healthmonitors list is empty: %v", elb.Healthmonitors)
		}

		if elb.Healthmonitors[0].OperatingStatus == "ONLINE" {
			return true, nil
		}

		if elb.Healthmonitors[0].OperatingStatus != "ONLINE" {
			time.Sleep(5 * time.Second)
			return false, nil
		}
		return false, fmt.Errorf("elb listener healthmonitor operating status: %v", elb.Healthmonitors[0].OperatingStatus)
	})
}
