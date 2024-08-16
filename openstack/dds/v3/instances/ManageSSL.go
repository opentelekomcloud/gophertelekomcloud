package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type SSLOpt struct {
	SSL        string `json:"ssl_option" required:"true"`
	InstanceId string `json:"-"`
}

func SwitchSSL(client *golangsdk.ServiceClient, opts SSLOpt) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "switch-ssl"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return ExtractJob(err, raw)
}
