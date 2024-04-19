package function

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type ExportOpts struct {
	Config string `q:"config"`
	Code   string `q:"code"`
	Type   string `q:"type"`
}

func Export(client *golangsdk.ServiceClient, funcURN string, opts ExportOpts) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("fgs", "functions", funcURN, "export").WithQueryParams(&opts).Build()
	if err != nil {
		return err
	}

	_, err = client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
