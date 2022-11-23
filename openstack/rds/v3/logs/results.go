package logs

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ErrorLogResult struct {
	golangsdk.Result
}

type ErrorLogResp struct {
	//
	ErrorLogList []Errorlog `json:"error_log_list"`
	//
	TotalRecord int `json:"total_record"`
}

type Errorlog struct {
	//
	Time string `json:"time"`
	//
	Level string `json:"level"`
	//
	Content string `json:"content"`
}

type ErrorLogPage struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if a ListResult contains no services.
func (r ErrorLogPage) IsEmpty() (bool, error) {
	data, err := ExtractErrorLog(r)
	if err != nil {
		return false, err
	}
	return len(data.ErrorLogList) == 0, err
}

func ExtractErrorLog(r pagination.Page) (ErrorLogResp, error) {
	var s ErrorLogResp
	err := (r.(ErrorLogPage)).ExtractInto(&s)
	return s, err
}

type SlowLogResp struct {
	//
	Slowloglist []Slowloglist `json:"slow_log_list"`
	//
	TotalRecord int `json:"total_record"`
}

type Slowloglist struct {
	//
	Count string `json:"count"`
	//
	Time string `json:"time"`
	//
	Locktime string `json:"lock_time"`
	//
	Rowssent string `json:"rows_sent"`
	//
	Rowsexamined string `json:"rows_examined"`
	//
	Database string `json:"database"`
	//
	Users string `json:"users"`
	//
	QuerySample string `json:"query_sample"`
	//
	Type string `json:"type"`
}

type SlowLogPage struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if a ListResult contains no services.
func (r SlowLogPage) IsEmpty() (bool, error) {
	data, err := ExtractSlowLog(r)
	if err != nil {
		return false, err
	}
	return len(data.Slowloglist) == 0, err
}

// ExtractCloudServers is a function that takes a ListResult and returns the services' information.
func ExtractSlowLog(r pagination.Page) (SlowLogResp, error) {
	var s SlowLogResp
	err := (r.(SlowLogPage)).ExtractInto(&s)
	return s, err
}

type UpdateConfigurationResponse struct {
	RestartRequired bool `json:"restart_required"`
}

type UpdateConfigurationResult struct {
	golangsdk.Result
}

func (r UpdateConfigurationResult) Extract() (*UpdateConfigurationResponse, error) {
	restartRequired := new(UpdateConfigurationResponse)
	err := r.ExtractInto(restartRequired)
	if err != nil {
		return nil, err
	}
	return restartRequired, nil
}
