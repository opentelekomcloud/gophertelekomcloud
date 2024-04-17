package alias

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateAliasOpts struct {
	FuncUrn                   string                    `json:"-"`
	AliasName                 string                    `json:"-"`
	Version                   string                    `json:"version" required:"true"`
	Description               string                    `json:"description,omitempty"`
	AdditionalVersionWeights  map[string]int            `json:"additional_version_weights,omitempty"`
	AdditionalVersionStrategy map[string]VectorStrategy `json:"additional_version_strategy,omitempty"`
}

func UpdateAlias(client *golangsdk.ServiceClient, opts UpdateAliasOpts) (*FuncAliasesResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("fgs", "functions", opts.FuncUrn, "aliases", opts.AliasName), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res FuncAliasesResp
	return &res, extract.Into(raw.Body, &res)
}
