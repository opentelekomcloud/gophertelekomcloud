package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// DisableBaseLogs will disable log switch.
func DisableBaseLogs(client *golangsdk.ServiceClient, clusterID string) error {
	action := "base_log_collect"
	return disableLogs(client, clusterID, &action)
}

// DisableRealTimeLogs will disable log ingestion.
func DisableRealTimeLogs(client *golangsdk.ServiceClient, clusterID string) error {
	action := "real_time_log_collect"
	return disableLogs(client, clusterID, &action)
}

// DisableLogs function will disable the log option for a CSS cluster.
func disableLogs(client *golangsdk.ServiceClient, clusterID string, action *string) error {
	url := client.ServiceURL("clusters", clusterID, "logs", "close")

	if action != nil && *action != "" {
		url += "?action=" + *action
	}

	_, err := client.Put(url, nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
