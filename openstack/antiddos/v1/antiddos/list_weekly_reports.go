package antiddos

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListWeeklyReports(client *golangsdk.ServiceClient, periodStartDate int64) (*ListWeeklyReportsResponse, error) {
	// GET /v1/{project_id}/antiddos/weekly
	url := client.ServiceURL("antiddos", "weekly") + "?period_start_date=" + strconv.FormatInt(periodStartDate, 10)
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListWeeklyReportsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListWeeklyReportsResponse struct {
	// Number of DDoS attacks intercepted in a week
	DDosInterceptTimes int `json:"ddos_intercept_times"`
	// Number of DDoS attacks intercepted in a week
	Weekdata []WeekData `json:"weekdata"`
	// Top 10 attacked IP addresses
	Top10 []Top10 `json:"top10"`
}

type Top10 struct {
	// EIP
	FloatingIpAddress string `json:"floating_ip_address"`
	// Number of DDoS attacks intercepted, including cleaning operations and black-holes
	Times int `json:"times"`
}

type WeekData struct {
	// Number of DDoS attacks intercepted
	DDosInterceptTimes int `json:"ddos_intercept_times"`
	// Number of DDoS blackholes
	DDosBlackholeTimes int `json:"ddos_blackhole_times"`
	// Maximum attack traffic
	MaxAttackBps int `json:"max_attack_bps"`
	// Maximum number of attack connections
	MaxAttackConns int `json:"max_attack_conns"`
	// Start date
	PeriodStartDate int64 `json:"period_start_date"`
}
