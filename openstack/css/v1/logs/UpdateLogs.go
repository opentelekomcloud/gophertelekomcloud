package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateLogConfigurationOpts struct {
	// These parameters are passed to the logs.UpdateLogs function.
	// Agency is the agency name used for the css cluster.
	Agency string `json:"agency" required:"true"`
	// BasePath is the obs path where the logs should be stored for the css cluster.
	BasePath string `json:"logBasePath" required:"true"`
	// Bucket is the obs bucket name to store the logs for the css cluster.
	Bucket string `json:"logBucket" required:"true"`
}

// UpdateLogs will change the cluster logging configurations based on UpdateLogConfigurationOpts.
func UpdateLogs(client *golangsdk.ServiceClient, clusterID string, opts UpdateLogConfigurationOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("clusters", clusterID, "logs", "settings"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
