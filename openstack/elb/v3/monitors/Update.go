package monitors

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Delay          int    `json:"delay,omitempty"`
	Timeout        int    `json:"timeout,omitempty"`
	MaxRetries     int    `json:"max_retries,omitempty"`
	MaxRetriesDown int    `json:"max_retries_down,omitempty"`
	URLPath        string `json:"url_path,omitempty"`
	DomainName     string `json:"domain_name,omitempty"`
	HTTPMethod     string `json:"http_method,omitempty"`
	ExpectedCodes  string `json:"expected_codes,omitempty"`
	Name           string `json:"name,omitempty"`
	AdminStateUp   *bool  `json:"admin_state_up,omitempty"`
	MonitorPort    int    `json:"monitor_port,omitempty"`
	Type           string `json:"type,omitempty"`

	// These fields are documented but marked unsupported by the reviewed OTC API.
	HTTPVersion        *string `json:"http_version,omitempty"`
	ContentCheckEnable *bool   `json:"content_check_enable,omitempty"`
	SendContent        *string `json:"send_content,omitempty"`
	ReceiveContent     *string `json:"receive_content,omitempty"`
	ContentFormat      string  `json:"content_format,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Monitor, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("healthmonitors", id).Build()
	if err != nil {
		return nil, err
	}
	b, err := build.RequestBody(opts, "healthmonitor")
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
	if err != nil {
		return nil, err
	}
	var res struct {
		Monitor Monitor `json:"healthmonitor"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.Monitor, nil
}
