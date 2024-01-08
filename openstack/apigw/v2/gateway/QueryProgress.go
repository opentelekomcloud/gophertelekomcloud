package gateway

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func QueryProgress(client *golangsdk.ServiceClient, id string) (*Progress, error) {
	raw, err := client.Get(client.ServiceURL("apigw", "instances", id, "progress"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Progress

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Progress struct {
	Progress  int    `json:"progress"`
	Status    string `json:"status"`
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}
