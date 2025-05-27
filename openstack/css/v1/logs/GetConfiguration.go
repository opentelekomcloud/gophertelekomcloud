package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type LogConfiguration struct {
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

// GetConfiguration function will query the details of CSS cluster logging and returns a LogConfiguration object.
func GetConfiguration(client *golangsdk.ServiceClient, clusterID string) (*LogConfiguration, error) {
	raw, err := client.Get(client.ServiceURL("clusters", clusterID, "logs", "settings"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res LogConfiguration
	err = extract.IntoStructPtr(raw.Body, &res, "logConfiguration")
	return &res, err
}
