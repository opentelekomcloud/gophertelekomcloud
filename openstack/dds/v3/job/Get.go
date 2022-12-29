package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*JobDDSInstance, error) {
	raw, err := client.Get(client.ServiceURL("jobs?id="+id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res JobDDSInstance
	err = extract.IntoStructPtr(raw.Body, &res, "job")
	return &res, err
}

type JobDDSInstance struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Status     string   `json:"status"`
	Created    string   `json:"created"`
	Ended      string   `json:"ended"`
	Progress   string   `json:"progress"`
	FailReason string   `json:"fail_reason"`
	Instance   Instance `json:"instance"`
}

type Instance struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
