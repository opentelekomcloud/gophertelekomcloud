package kibana

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

func Disable(client *golangsdk.ServiceClient, opts ManageOpts) error {
	request := Whitelist{
		EipSize: opts.EipSize,
		ElbWhiteList: KibanaPublicWhitelist{
			EnableWhitelist: opts.EnableWhiteList,
			WhiteList:       opts.Whitelist,
		},
	}

	b, err := build.RequestBody(request, "")
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL("clusters", opts.ClusterId, "publickibana", "close"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
