package instances

import "github.com/opentelekomcloud/gophertelekomcloud"

type SpecCode struct {
	Speccode string `json:"spec_code" required:"true"`
}

type ResizeFlavorOpts struct {
	ResizeFlavor *SpecCode `json:"resize_flavor" required:"true"`
}

type ResizeFlavorBuilder interface {
	ResizeFlavorMap() (map[string]interface{}, error)
}

func (opts ResizeFlavorOpts) ResizeFlavorMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Resize(client *golangsdk.ServiceClient, opts ResizeFlavorBuilder, instanceId string) (r ResizeFlavorResult) {
	b, err := opts.ResizeFlavorMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(client.ServiceURL("instances", instanceId, "action"), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})

	return
}
