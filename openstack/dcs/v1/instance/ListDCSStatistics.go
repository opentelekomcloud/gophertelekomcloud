package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListDCSStatistics(client *golangsdk.ServiceClient) ([]InstanceStatistic, error) {
	raw, err := client.Get(client.ServiceURL("instances", "statistic"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []InstanceStatistic
	err = extract.IntoSlicePtr(raw.Body, &res, "statistics")
	return res, err
}

type InstanceStatistic struct {
	// Incoming traffic (kbit/s) of the DCS instance
	InputKbps string `json:"input_kbps"`
	// Outgoing traffic (kbit/s) of the DCS instance
	OutputKbps string `json:"output_kbps"`
	// DCS instance ID
	InstanceId string `json:"instance_id"`
	// Number of cached data records
	Keys int64 `json:"keys"`
	// Size of the used memory in MB
	UsedMemory int64 `json:"used_memory"`
	// Overall memory size in MB
	MaxMemory int64 `json:"max_memory"`
	// Number of times the GET command is run
	CmdGetCount int64 `json:"cmd_get_count"`
	// Number of times the SET command is run
	CmdSetCount int64 `json:"cmd_set_count"`
	// Percentage of CPU usage
	UsedCpu string `json:"used_cpu"`
}
