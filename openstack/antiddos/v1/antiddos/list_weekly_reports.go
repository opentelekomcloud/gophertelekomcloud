package antiddos

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListWeeklyReports(client *golangsdk.ServiceClient, periodStartDate time.Time) (*ListWeeklyReportsResponse, error) {
	// GET /v1/{project_id}/antiddos/weekly
	raw, err := client.Get(
		client.ServiceURL("antiddos", "weekly")+"?period_start_date="+strconv.FormatInt(periodStartDate.Unix()*1000, 10),
		nil, nil)
	if err != nil {
		return nil, err
	}

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
		// Number of DDoS attacks intercepted, including cleaning operations and black-holes
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

func (r *ListWeeklyReportsResponse) UnmarshalJSON(b []byte) error {
	type tmp ListWeeklyReportsResponse
	var s struct {
		tmp
		Weekdata []struct {
			// Number of DDoS attacks intercepted
			DDosInterceptTimes int `json:"ddos_intercept_times,"`
			// Number of DDoS blackholes
			DDosBlackholeTimes int `json:"ddos_blackhole_times,"`
			// Maximum attack traffic
			MaxAttackBps int `json:"max_attack_bps,"`
			// Maximum number of attack connections
			MaxAttackConns int `json:"max_attack_conns,"`
			// Start date
			PeriodStartDate int64 `json:"period_start_date,"`
		} `json:"weekdata,"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	*r = ListWeeklyReportsResponse(s.tmp)
	r.Weekdata = make([]WeekData, len(s.Weekdata))

	for idx, val := range s.Weekdata {
		r.Weekdata[idx] = WeekData{
			DDosInterceptTimes: val.DDosBlackholeTimes,
			DDosBlackholeTimes: val.DDosBlackholeTimes,
			MaxAttackBps:       val.MaxAttackBps,
			MaxAttackConns:     val.MaxAttackConns,
			PeriodStartDate:    time.Unix(val.PeriodStartDate/1000, 0).UTC(),
		}
	}

	return nil
}
