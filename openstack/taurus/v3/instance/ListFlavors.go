package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListFlavorsOpts struct {
	DatabaseName         string `json:"-"`
	VersionName          string `q:"version_name"`
	AvailabilityZoneMode string `q:"availability_zone_mode" required:"true"`
	SpecCode             string `q:"spec_code"`
}

func ListFlavors(client *golangsdk.ServiceClient, opts ListFlavorsOpts) ([]MysqlFlavorInfo, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("flavors", opts.DatabaseName).
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res struct {
		Flavors []MysqlFlavorInfo `json:"flavors"`
	}
	err = extract.IntoStructPtr(raw.Body, &res, "")
	return res.Flavors, err
}

type MysqlFlavorInfo struct {
	Id           string            `json:"id"`
	SpecCode     string            `json:"spec_code"`
	Vcpus        string            `json:"vcpus"`
	Ram          string            `json:"ram"`
	Type         string            `json:"type"`
	VersionName  string            `json:"version_name"`
	InstanceMode string            `json:"instance_mode"`
	AzStatus     map[string]string `json:"az_status"`
}
