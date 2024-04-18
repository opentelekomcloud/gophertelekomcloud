package util

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetFuncLts(client *golangsdk.ServiceClient, funcURN string) (*FuncLtsResp, error) {
	raw, err := client.Get(client.ServiceURL("fgs", "functions", funcURN, "lts-log-detail"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res FuncLtsResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type FuncLtsResp struct {
	GroupName  string `json:"group_name"`
	GroupId    string `json:"group_id"`
	StreamId   string `json:"stream_id"`
	StreamName string `json:"stream_name"`
}
