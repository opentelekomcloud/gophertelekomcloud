package antiddos

import (
	"strconv"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListWeeklyReports(client *golangsdk.ServiceClient, periodStartDate time.Time) (*ListWeeklyReportsResponse, error) {
	raw, err := client.Get(
		client.ServiceURL("antiddos", "weekly")+"?period_start_date="+strconv.FormatInt(periodStartDate.Unix()*1000, 10),
		nil, nil)

	var res ListWeeklyReportsResponse
	err = extract.Into(raw, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type ListWeeklyReportsResponse struct {
	// Number of DDoS attacks intercepted in a week
	DDosInterceptTimes int `json:"ddos_intercept_times,"`
	// Number of DDoS attacks intercepted in a week
	Weekdata []WeekData `json:"-"`
	// Top 10 attacked IP addresses
	Top10 []struct {
		// EIP
		FloatingIpAddress string `json:"floating_ip_address,"`
		// Number of DDoS attacks intercepted, including cleaning operations and blackholes
		Times int `json:"times,"`
	} `json:"top10,"`
}

type WeekData struct {
	// Number of DDoS attacks intercepted
	DDosInterceptTimes int `json:"ddos_intercept_times,"`
	// Number of DDoS blackholes
	DDosBlackholeTimes int `json:"ddos_blackhole_times,"`
	// Maximum attack traffic
	MaxAttackBps int `json:"max_attack_bps,"`
	// Maximum number of attack connections
	MaxAttackConns int `json:"max_attack_conns,"`
	// Start date
	PeriodStartDate time.Time `json:"period_start_date,"`
}
