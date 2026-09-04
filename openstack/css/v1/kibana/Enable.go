package kibana

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ManageOpts struct {
	ClusterId       string `json:"clusterId"`
	EipSize         int    `json:"-"`
	EnableWhiteList bool   `json:"-"`
	Whitelist       string `json:"-"`
}

func Enable(client *golangsdk.ServiceClient, opts ManageOpts) error {
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

	_, err = client.Post(client.ServiceURL("clusters", opts.ClusterId, "publickibana", "open"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}

type Whitelist struct {
	EipSize      int                   `json:"eipSize"`
	ElbWhiteList KibanaPublicWhitelist `json:"elbWhiteList"`
}

type KibanaPublicWhitelist struct {
	EnableWhitelist bool   `json:"enableWhiteList"`
	WhiteList       string `json:"whiteList"`
}
