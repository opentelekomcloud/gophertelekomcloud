package volumeactions

import "github.com/opentelekomcloud/gophertelekomcloud"

type TerminateConnectionOptsBuilder interface {
	ToVolumeTerminateConnectionMap() (map[string]interface{}, error)
}

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

func TerminateConnection(client *golangsdk.ServiceClient, id string, opts TerminateConnectionOptsBuilder) (r TerminateConnectionResult) {
	b, err := opts.ToVolumeTerminateConnectionMap()
	if err != nil {
		r.Err = err
		return
	}
	raw, err := client.Post(client.ServiceURL("volumes", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
