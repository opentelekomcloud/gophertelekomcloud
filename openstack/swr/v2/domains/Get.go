package domains

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, opts GetOpts) (*AccessDomain, error) {
	// GET /v2/manage/namespaces/{namespace}/repositories/{repository}/access-domains/{access_domain}
	url := fmt.Sprintf("%s/%s", client.ServiceURL("manage", "namespaces", opts.Namespace, "repos", opts.Repository, "access-domains"), opts.AccessDomain)
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res AccessDomain
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type AccessDomain struct {
	Exist bool `json:"exist"`
	// Organization name.
	Namespace string `json:"namespace"`
	// Image repository name.
	Repository string `json:"repository"`
	// Name of the account used for image sharing
	AccessDomain string `json:"access_domain"`
	// Permission
	Permit string `json:"permit"`
	// Expiration time.
	Deadline string `json:"deadline"`
	// Description
	Description string `json:"description"`
	// Creator ID.
	CreatorID string `json:"creator_id"`
	// Name of the creator.
	CreatorName string `json:"creator_name"`
	// Time when an image is created. It is the UTC standard time.
	Created string `json:"created"`
	// Time when an image is updated. It is the UTC standard time.
	Updated string `json:"updated"`
	// Status. `true`: valid `false`: expired
	Status bool `json:"status"`
}
