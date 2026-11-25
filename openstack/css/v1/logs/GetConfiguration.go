package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BaseLogConfiguration struct {
	// These parameters are passed to the logs.GetLogConfiguration function.
	// Log backup ID.
	ID string `json:"id"`
	// CSS cluster ID.
	ClusterID string `json:"clusterId"`
	// The bucket where the logs should be stored.
	ObsBucket string `json:"obsBucket"`
	// The agency name.
	Agency string `json:"agency"`
	// Update time.
	UpdateAt int `json:"updateAt"`
	// Storage path of backup logs in the OBS bucket.
	BasePath string `json:"basePath"`
	// Indicates whether to enable automatic backup.
	AutoEnable bool `json:"autoEnable"`
	// Start time of automatic log backup.
	Period string `json:"period"`
	// Indicates whether to enable the log function.
	LogSwitch bool `json:"logSwitch"`
}

type RealTimeLogConfiguration struct {
	// Log backup ID.
	ID string `json:"id"`
	// CSS cluster ID.
	ClusterID string `json:"clusterId"`
	// Prefix of the index for saving logs.
	IndexPrefix string `json:"indexPrefix"`
	// Log retention duration.
	KeepDays int `json:"keepDays"`
	// ID of the target cluster where logs are saved.
	TargetClusterId string `json:"targetClusterId"`
	// Status of a real-time log collection task.
	Status string `json:"status"`
	// Start time of a real-time log collection task.
	CreateAt int64 `json:"createAt"`
	// Update time.
	UpdateAt int `json:"updateAt"`
}

// GetBaseLogConfiguration will return the base log collect configurations.
func GetBaseLogConfiguration(client *golangsdk.ServiceClient, clusterID string) (interface{}, error) {
	action := "base_log_collect"
	return GetConfiguration(client, clusterID, &action)
}

// GetRealTimeLogConfiguration will return the real time log collect configurations.
func GetRealTimeLogConfiguration(client *golangsdk.ServiceClient, clusterID string) (interface{}, error) {
	action := "real_time_log_collect"
	return GetConfiguration(client, clusterID, &action)
}

// GetConfiguration function will query the details of CSS cluster logging and returns a LogConfiguration object.
func GetConfiguration(client *golangsdk.ServiceClient, clusterID string, action *string) (interface{}, error) {
	url := client.ServiceURL("clusters", clusterID, "logs", "settings")

	if action != nil && *action != "" {
		url += "?action=" + *action
	}

	println(url)

	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	record := ""

	if *action == "real_time_log_collect" {
		var res RealTimeLogConfiguration
		record = "realTimeLogCollectRecord"
		err = extract.IntoStructPtr(raw.Body, &res, record)
		return &res, err
	} else if *action == "base_log_collect" || *action == "" {
		var res BaseLogConfiguration
		record = "logConfiguration"
		err = extract.IntoStructPtr(raw.Body, &res, record)
		return &res, err
	}

	return nil, err

}
