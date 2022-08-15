package antiddos

import (
	"encoding/json"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type commonResult struct {
	golangsdk.Result
}

type CreateResult struct {
	commonResult
}

func (r CreateResult) Extract() (*CreateResponse, error) {
	var response CreateResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type CreateResponse struct {
	// Internal error code
	ErrorCode string `json:"error_code,"`

	// Internal error description
	ErrorDescription string `json:"error_description,"`

	// ID of a task. This ID can be used to query the status of the task. This field is reserved for use in task auditing later. It is temporarily unused.
	TaskId string `json:"task_id,"`
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

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*GetResponse, error) {
	var response GetResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type GetResponse struct {
	// Whether L7 defense has been enabled
	EnableL7 bool `json:"enable_L7,"`

	// Position ID of traffic. The value ranges from 1 to 9.
	TrafficPosId int `json:"traffic_pos_id,"`

	// Position ID of number of HTTP requests. The value ranges from 1 to 15.
	HttpRequestPosId int `json:"http_request_pos_id,"`

	// Position ID of access limit during cleaning. The value ranges from 1 to 8.
	CleaningAccessPosId int `json:"cleaning_access_pos_id,"`

	// Application type ID. Possible values: 0 1
	AppTypeId int `json:"app_type_id,"`
}

type GetStatusResult struct {
	commonResult
}

func (r GetStatusResult) Extract() (*GetStatusResponse, error) {
	var response GetStatusResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type GetStatusResponse struct {
	// Defense status
	Status string `json:"status,"`
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

type UpdateResult struct {
	commonResult
}

func (r UpdateResult) Extract() (*UpdateResponse, error) {
	var response UpdateResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type UpdateResponse struct {
	// Internal error code
	ErrorCode string `json:"error_code,"`

	// Internal error description
	ErrorDescription string `json:"error_description,"`

	// ID of a task. This ID can be used to query the status of the task. This field is reserved for use in task auditing later. It is temporarily unused.
	TaskId string `json:"task_id,"`
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
