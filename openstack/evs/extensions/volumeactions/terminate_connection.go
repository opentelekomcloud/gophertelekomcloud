package volumeactions

import "github.com/opentelekomcloud/gophertelekomcloud"

type TerminateConnectionOpts struct {
	IP        string   `json:"ip,omitempty"`
	Host      string   `json:"host,omitempty"`
	Initiator string   `json:"initiator,omitempty"`
	Wwpns     []string `json:"wwpns,omitempty"`
	Wwnns     string   `json:"wwnns,omitempty"`
	Multipath *bool    `json:"multipath,omitempty"`
	Platform  string   `json:"platform,omitempty"`
	OSType    string   `json:"os_type,omitempty"`
}

func (opts TerminateConnectionOpts) ToVolumeTerminateConnectionMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "connector")
	return map[string]interface{}{"os-terminate_connection": b}, err
}

func TerminateConnection(client *golangsdk.ServiceClient, id string, opts TerminateConnectionOpts) (err error) {
	b, err := golangsdk.BuildRequestBody(opts, "connector")
	b = map[string]interface{}{"os-terminate_connection": b}
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("volumes", id, "action"), b, nil, nil)
	return
}
