package alias

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListAlias(client *golangsdk.ServiceClient, funcURN string) ([]FuncAliases, error) {
	raw, err := client.Get(client.ServiceURL("fgs", "functions", funcURN, "aliases"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []FuncAliases
	err = extract.IntoSlicePtr(raw.Body, &res, "")
	return res, err
}

type FuncAliases struct {
	Name                     string         `json:"name"`
	Version                  string         `json:"version"`
	Description              string         `json:"description"`
	LastModified             string         `json:"last_modified"`
	AliasUrn                 string         `json:"alias_urn"`
	AdditionalVersionWeights map[string]int `json:"additional_version_weights"`
}
