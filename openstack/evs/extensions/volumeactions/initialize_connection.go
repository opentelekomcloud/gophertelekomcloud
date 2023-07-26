package volumeactions

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type InitializeConnectionOpts struct {
	IP        string   `json:"ip,omitempty"`
	Host      string   `json:"host,omitempty"`
	Initiator string   `json:"initiator,omitempty"`
	Wwpns     []string `json:"wwpns,omitempty"`
	Wwnns     string   `json:"wwnns,omitempty"`
	Multipath *bool    `json:"multipath,omitempty"`
	Platform  string   `json:"platform,omitempty"`
	OSType    string   `json:"os_type,omitempty"`
}

func InitializeConnection(client *golangsdk.ServiceClient, id string, opts InitializeConnectionOpts) (map[string]any, error) {
	b, err := golangsdk.BuildRequestBody(opts, "connector")
	b = map[string]any{"os-initialize_connection": b}
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("volumes", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		ConnectionInfo map[string]any `json:"connection_info"`
	}
	err = extract.Into(raw.Body, &res)
	return res.ConnectionInfo, err
}
