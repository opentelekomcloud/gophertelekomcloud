package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateLogConfigurationOpts struct {
	// These parameters are passed to the logs.UpdateLogs function.
	// Agency is the agency name used for the css cluster.
	Agency string `json:"agency,omitempty"`
	// BasePath is the obs path where the logs should be stored for the css cluster.
	BasePath string `json:"logBasePath,omitempty"`
	// Bucket is the obs bucket name to store the logs for the css cluster.
	Bucket string `json:"logBucket,omitempty"`
	// Index prefix for storing logs.
	IndexPrefix string `json:"index_prefix,omitempty"`
	// Log retention duration.
	KeepDays int `json:"keep_days,omitempty"`
	// Specifies the target cluster for saving logs.
	TargetClusterId string `json:"target_cluster_id,omitempty"`
}

// UpdateBaseLogs will change the base log collect configurations.
func UpdateBaseLogs(client *golangsdk.ServiceClient, clusterID string, opts UpdateLogConfigurationOpts) error {
	action := "base_log_collect"
	return updateLogs(client, clusterID, opts, &action)
}

// UpdateRealTimeLogs will change the real time log collect configurations.
func UpdateRealTimeLogs(client *golangsdk.ServiceClient, clusterID string, opts UpdateLogConfigurationOpts) error {
	action := "real_time_log_collect"
	return updateLogs(client, clusterID, opts, &action)
}

// UpdateLogs will change the cluster logging configurations based on UpdateLogConfigurationOpts.
func updateLogs(client *golangsdk.ServiceClient, clusterID string, opts UpdateLogConfigurationOpts, action *string) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	url := client.ServiceURL("clusters", clusterID, "logs", "settings")

	if action != nil && *action != "" {
		url += "?action=" + *action
	}

	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
