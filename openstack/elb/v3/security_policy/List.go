package security_policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Marker      string   `q:"marker"`
	Limit       string   `q:"limit"`
	PageReverse bool     `q:"page_reverse"`
	ID          []string `q:"id"`
	Name        []string `q:"name"`
	Description []string `q:"description"`
	Protocols   []string `q:"protocols"`
	Ciphers     []string `q:"ciphers"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]PolicyRef, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("security-policies")+q.String(), nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
	if err != nil {
		return nil, err
	}

	var res []PolicyRef

	err = extract.IntoSlicePtr(raw.Body, &res, "security_policies")
	return res, err

}
