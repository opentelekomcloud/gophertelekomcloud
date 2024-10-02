package antiddos

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListDailyReport(client *golangsdk.ServiceClient, floatingIpId string) ([]Data, error) {
	// GET /v1/{project_id}/antiddos/{floating_ip_id}/daily
	raw, err := client.Get(client.ServiceURL("antiddos", floatingIpId, "daily"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Data
	err = extract.IntoSlicePtr(raw.Body, &res, "data")
	return res, err
}

type Data struct {
	// Start time
	PeriodStart int `json:"period_start"`
	// Inbound traffic (bit/s)
	BpsIn int `json:"bps_in"`
	// Attack traffic (bit/s)
	BpsAttack int `json:"bps_attack"`
	// Total traffic
	TotalBps int `json:"total_bps"`
	// Inbound packet rate (number of packets per second)
	PpsIn int `json:"pps_in"`
	// Attack packet rate (number of packets per second)
	PpsAttack int `json:"pps_attack"`
	// Total packet rate
	TotalPps int `json:"total_pps"`
}
