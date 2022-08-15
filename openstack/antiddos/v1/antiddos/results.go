package antiddos

import (
	"encoding/json"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type commonResult struct {
	golangsdk.Result
}

type DeleteResult struct {
	commonResult
}

func (r DeleteResult) Extract() (*DeleteResponse, error) {
	var response DeleteResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type DeleteResponse struct {
	// Internal error code
	ErrorCode string `json:"error_code,"`

	// Internal error description
	ErrorDescription string `json:"error_description,"`

	// ID of a task. This ID can be used to query the status of the task. This field is reserved for use in task auditing later. It is temporarily unused.
	TaskId string `json:"task_id,"`
}

type GetTaskResult struct {
	commonResult
}

func (r GetTaskResult) Extract() (*GetTaskResponse, error) {
	var response GetTaskResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type GetTaskResponse struct {
	// Status of a task, which can be one of the following: success, failed, waiting, running, preprocess, ready
	TaskStatus string `json:"task_status,"`

	// Additional information about a task
	TaskMsg string `json:"task_msg,"`
}

func (r *ListWeeklyReportsResponse) UnmarshalJSON(b []byte) error {
	type tmp ListWeeklyReportsResponse
	var s struct {
		tmp
		Weekdata []struct {
			// Number of DDoS attacks intercepted
			DdosInterceptTimes int `json:"ddos_intercept_times,"`

			// Number of DDoS blackholes
			DdosBlackholeTimes int `json:"ddos_blackhole_times,"`

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
			DDosInterceptTimes: val.DdosBlackholeTimes,
			DDosBlackholeTimes: val.DdosBlackholeTimes,
			MaxAttackBps:       val.MaxAttackBps,
			MaxAttackConns:     val.MaxAttackConns,
			PeriodStartDate:    time.Unix(val.PeriodStartDate/1000, 0).UTC(),
		}
	}

	return nil
}
