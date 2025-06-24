package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type EnableLogsOpts struct {
	// These parameters are passed to the logs.EnableLogs function.
	// Agency is the agency name used for the css cluster.
	Agency string `json:"agency" required:"true"`
	// BasePath is the obs path where the logs should be stored for the css cluster.
	BasePath string `json:"logBasePath" required:"true"`
	// Bucket is the obs bucket name to store the logs for the css cluster.
	Bucket string `json:"logBucket" required:"true"`
}

// EnableLogs function is used to enable the log switch of a CSS cluster base on EnableLogsOpts.
func EnableLogs(client *golangsdk.ServiceClient, clusterID string, opts EnableLogsOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	url := client.ServiceURL("clusters", clusterID, "logs", "open")

	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
