package v3

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

func WaitForGaussJob(client *golangsdk.ServiceClient, jobId string, timeout int) (*GetJobInfoDetail, error) {
	var res *GetJobInfoDetail

	err := golangsdk.WaitFor(timeout, func() (bool, error) {
		cur, err := ShowJobInfo(client, jobId)
		if err != nil {
			return false, err
		}
		res = cur

		if cur.Status == "Completed" {
			return true, nil
		}

		if cur.Status == "Failed" {
			return false, fmt.Errorf("job %s failed: %s", jobId, cur.FailReason)
		}

		return false, nil
	})

	return res, err
}
