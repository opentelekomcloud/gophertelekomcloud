package volumeactions

import "github.com/opentelekomcloud/gophertelekomcloud"

type InitializeConnectionOptsBuilder interface {
	ToVolumeInitializeConnectionMap() (map[string]interface{}, error)
}

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

func (opts InitializeConnectionOpts) ToVolumeInitializeConnectionMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "connector")
	return map[string]interface{}{"os-initialize_connection": b}, err
}

func InitializeConnection(client *golangsdk.ServiceClient, id string, opts InitializeConnectionOptsBuilder) (r InitializeConnectionResult) {
	b, err := opts.ToVolumeInitializeConnectionMap()
	if err != nil {
		r.Err = err
		return
	}
	raw, err := client.Post(client.ServiceURL("volumes", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	return
}
