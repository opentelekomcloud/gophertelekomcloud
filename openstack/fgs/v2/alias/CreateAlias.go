package alias

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateAliasOpts struct {
	FuncUrn                   string                    `json:"-"`
	Name                      string                    `json:"name" required:"true"`
	Version                   string                    `json:"version" required:"true"`
	Description               string                    `json:"description,omitempty"`
	AdditionalVersionWeights  map[string]int            `json:"additional_version_weights,omitempty"`
	AdditionalVersionStrategy map[string]VectorStrategy `json:"additional_version_strategy,omitempty"`
}

type VectorStrategy struct {
	CombineType string                `json:"combine_type"`
	Rules       *VersionStrategyRules `json:"rules"`
}

type VersionStrategyRules struct {
	RuleType string `json:"rule_type"`
	Param    string `json:"param"`
	Op       string `json:"op"`
	Value    string `json:"value"`
}

func CreateAlias(client *golangsdk.ServiceClient, opts CreateAliasOpts) (*FuncAliasesResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("fgs", "functions", opts.FuncUrn, "aliases"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res FuncAliasesResp
	return &res, extract.Into(raw.Body, &res)
}

type FuncAliasesResp struct {
	Name                      string                    `json:"name"`
	Version                   string                    `json:"version"`
	Description               string                    `json:"description"`
	LastModified              string                    `json:"last_modified"`
	AliasUrn                  string                    `json:"alias_urn"`
	AdditionalVersionWeights  map[string]int            `json:"additional_version_weights"`
	AdditionalVersionStrategy map[string]VectorStrategy `json:"additional_version_strategy"`
}
