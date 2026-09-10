package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Category       string `q:"category"`
	EnableStatus   *bool  `q:"enablestatus"`
	ClusterGroupID string `q:"clustergroupid"`
	Limit          int    `q:"limit"`
	Offset         int    `q:"offset"`
	OrderBy        string `q:"order_by"`
	Order          string `q:"order"`
	ManageType     string `q:"managetype"`
	ClusterIDs     string `q:"clusterids"`
}

type ListResponse struct {
	Items []Cluster `json:"items"`
	Total int       `json:"total"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("clusters").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	return &res, extract.Into(raw.Body, &res)
}
